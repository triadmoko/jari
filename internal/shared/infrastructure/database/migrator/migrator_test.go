package migrator

import (
	"database/sql"
	"errors"
	"testing"
	"testing/fstest"

	"app/migrations"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	queryTable   = `SELECT to_regclass($1)::text`
	queryVersion = `SELECT version, dirty FROM schema_migrations LIMIT 1`
)

func migrationFS(names ...string) fstest.MapFS {
	fsys := fstest.MapFS{}
	for _, n := range names {
		fsys[n] = &fstest.MapFile{Data: []byte("SELECT 1;")}
	}
	return fsys
}

func TestLatestVersion(t *testing.T) {
	tests := []struct {
		name string
		fsys fstest.MapFS
		want uint
	}{
		{"no migrations", migrationFS(), 0},
		{"single", migrationFS("001_a.up.sql", "001_a.down.sql"), 1},
		{"multiple unordered", migrationFS("010_c.up.sql", "002_b.up.sql", "001_a.up.sql"), 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LatestVersion(tt.fsys)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLatestVersion_EmbeddedMigrations(t *testing.T) {
	got, err := LatestVersion(migrations.FS)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, got, uint(1), "embedded migrations must include 001_create_users_table")
}

func newMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })
	return db, mock
}

func expectTable(mock sqlmock.Sqlmock, exists bool) {
	var value any
	if exists {
		value = MigrationsTable
	}
	mock.ExpectQuery(queryTable).WithArgs(MigrationsTable).
		WillReturnRows(sqlmock.NewRows([]string{"to_regclass"}).AddRow(value))
}

func expectVersion(mock sqlmock.Sqlmock, version int64, dirty bool) {
	mock.ExpectQuery(queryVersion).
		WillReturnRows(sqlmock.NewRows([]string{"version", "dirty"}).AddRow(version, dirty))
}

func TestCheckSchema(t *testing.T) {
	fsys := migrationFS("001_a.up.sql", "002_b.up.sql")

	tests := []struct {
		name    string
		setup   func(sqlmock.Sqlmock)
		wantErr *SchemaError
	}{
		{
			name:    "migrations table missing",
			setup:   func(m sqlmock.Sqlmock) { expectTable(m, false) },
			wantErr: &SchemaError{Current: 0, Expected: 2},
		},
		{
			name: "table exists but empty",
			setup: func(m sqlmock.Sqlmock) {
				expectTable(m, true)
				m.ExpectQuery(queryVersion).WillReturnRows(sqlmock.NewRows([]string{"version", "dirty"}))
			},
			wantErr: &SchemaError{Current: 0, Expected: 2},
		},
		{
			name: "all rolled back (nil version)",
			setup: func(m sqlmock.Sqlmock) {
				expectTable(m, true)
				expectVersion(m, -1, false)
			},
			wantErr: &SchemaError{Current: 0, Expected: 2},
		},
		{
			name: "pending",
			setup: func(m sqlmock.Sqlmock) {
				expectTable(m, true)
				expectVersion(m, 1, false)
			},
			wantErr: &SchemaError{Current: 1, Expected: 2},
		},
		{
			name: "dirty",
			setup: func(m sqlmock.Sqlmock) {
				expectTable(m, true)
				expectVersion(m, 2, true)
			},
			wantErr: &SchemaError{Current: 2, Expected: 2, Dirty: true},
		},
		{
			name: "database newer than binary",
			setup: func(m sqlmock.Sqlmock) {
				expectTable(m, true)
				expectVersion(m, 3, false)
			},
			wantErr: &SchemaError{Current: 3, Expected: 2},
		},
		{
			name: "up to date",
			setup: func(m sqlmock.Sqlmock) {
				expectTable(m, true)
				expectVersion(m, 2, false)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := newMock(t)
			tt.setup(mock)

			err := CheckSchema(db, fsys)

			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				var se *SchemaError
				require.True(t, errors.As(err, &se), "expected *SchemaError, got %v", err)
				assert.Equal(t, tt.wantErr, se)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCheckSchema_QueryError(t *testing.T) {
	db, mock := newMock(t)
	mock.ExpectQuery(queryTable).WithArgs(MigrationsTable).WillReturnError(errors.New("connection refused"))

	err := CheckSchema(db, migrationFS("001_a.up.sql"))

	require.Error(t, err)
	var se *SchemaError
	assert.False(t, errors.As(err, &se))
	assert.Contains(t, err.Error(), "connection refused")
}

func TestSchemaError_Message(t *testing.T) {
	assert.Contains(t, (&SchemaError{Current: 0, Expected: 1}).Error(), "current version 0, expected 1")
	assert.Contains(t, (&SchemaError{Current: 1, Expected: 1, Dirty: true}).Error(), "dirty at version 1")
	assert.Contains(t, (&SchemaError{Current: 2, Expected: 1}).Error(), "newer than this binary")
}

func TestValidateArgs(t *testing.T) {
	valid := [][]string{
		{"up"}, {"down"}, {"down", "2"}, {"down", "all"}, {"version"}, {"force", "1"}, {"force", "-1"},
	}
	for _, args := range valid {
		assert.NoError(t, ValidateArgs(args), "%v", args)
	}

	invalid := [][]string{
		nil, {"sideways"}, {"up", "1"}, {"down", "0"}, {"down", "x"}, {"down", "1", "2"},
		{"version", "x"}, {"force"}, {"force", "a"}, {"force", "-2"},
	}
	for _, args := range invalid {
		assert.ErrorIs(t, ValidateArgs(args), ErrUsage, "%v", args)
	}
}
