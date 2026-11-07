// Package migrator applies and verifies database migrations using golang-migrate
// with migration files provided as an fs.FS (normally the embedded migrations.FS).
package migrator

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// MigrationsTable is the table golang-migrate uses to track the schema version.
const MigrationsTable = "schema_migrations"

const Usage = `usage: migrate <command>
  up              apply all pending migrations
  down [N|all]    roll back N migrations (default 1) or all of them
  version         print the current schema version
  force V         set the schema version to V and clear the dirty flag`

var ErrUsage = errors.New(Usage)

// SchemaError reports that the database schema does not match the migrations
// bundled with the binary.
type SchemaError struct {
	Current  uint
	Expected uint
	Dirty    bool
}

func (e *SchemaError) Error() string {
	switch {
	case e.Dirty:
		return fmt.Sprintf("database schema is dirty at version %d (expected %d): a previous migration failed; "+
			"fix the schema manually, then run `migrate force <version>` and `migrate up`", e.Current, e.Expected)
	case e.Current < e.Expected:
		return fmt.Sprintf("database schema has pending migrations: current version %d, expected %d; run `migrate up`",
			e.Current, e.Expected)
	default:
		return fmt.Sprintf("database schema is newer than this binary: current version %d, expected %d; "+
			"deploy a newer build or roll back with `migrate down`", e.Current, e.Expected)
	}
}

// LatestVersion returns the highest migration version found in fsys, or 0 when
// fsys contains no migrations.
func LatestVersion(fsys fs.FS) (uint, error) {
	src, err := iofs.New(fsys, ".")
	if err != nil {
		return 0, fmt.Errorf("read migrations: %w", err)
	}
	defer src.Close()

	version, err := src.First()
	if errors.Is(err, fs.ErrNotExist) {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("read migrations: %w", err)
	}
	for {
		next, err := src.Next(version)
		if errors.Is(err, fs.ErrNotExist) {
			return version, nil
		}
		if err != nil {
			return 0, fmt.Errorf("read migrations: %w", err)
		}
		version = next
	}
}

// CurrentVersion reads the schema version recorded by golang-migrate. A missing
// migrations table means no migration has run yet (version 0).
func CurrentVersion(db *sql.DB) (version uint, dirty bool, err error) {
	var table sql.NullString
	if err := db.QueryRow(`SELECT to_regclass($1)::text`, MigrationsTable).Scan(&table); err != nil {
		return 0, false, fmt.Errorf("check migrations table: %w", err)
	}
	if !table.Valid {
		return 0, false, nil
	}

	var v int64
	err = db.QueryRow(`SELECT version, dirty FROM `+MigrationsTable+` LIMIT 1`).Scan(&v, &dirty)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("read schema version: %w", err)
	}
	if v < 0 { // golang-migrate stores -1 (NilVersion) when every migration was rolled back
		return 0, dirty, nil
	}
	return uint(v), dirty, nil
}

// CheckSchema verifies that the database is migrated exactly to the latest
// version in fsys and is not dirty. It returns a *SchemaError on mismatch.
func CheckSchema(db *sql.DB, fsys fs.FS) error {
	expected, err := LatestVersion(fsys)
	if err != nil {
		return err
	}
	current, dirty, err := CurrentVersion(db)
	if err != nil {
		return err
	}
	if dirty || current != expected {
		return &SchemaError{Current: current, Expected: expected, Dirty: dirty}
	}
	return nil
}

type command struct {
	name  string
	steps int  // down: number of steps
	all   bool // down: roll back everything
	force int  // force: target version
}

func parseArgs(args []string) (command, error) {
	if len(args) == 0 {
		return command{}, ErrUsage
	}
	cmd := command{name: args[0]}
	rest := args[1:]
	switch cmd.name {
	case "up", "version":
		if len(rest) != 0 {
			return command{}, ErrUsage
		}
	case "down":
		cmd.steps = 1
		if len(rest) > 1 {
			return command{}, ErrUsage
		}
		if len(rest) == 1 {
			if rest[0] == "all" {
				cmd.all = true
				break
			}
			n, err := strconv.Atoi(rest[0])
			if err != nil || n < 1 {
				return command{}, ErrUsage
			}
			cmd.steps = n
		}
	case "force":
		if len(rest) != 1 {
			return command{}, ErrUsage
		}
		v, err := strconv.Atoi(rest[0])
		if err != nil || v < -1 {
			return command{}, ErrUsage
		}
		cmd.force = v
	default:
		return command{}, ErrUsage
	}
	return cmd, nil
}

// ValidateArgs checks the migrate arguments without touching the database.
func ValidateArgs(args []string) error {
	_, err := parseArgs(args)
	return err
}

// Run executes a migrate subcommand (see Usage) against db using the migrations
// in fsys. It takes ownership of db and closes it when done.
func Run(db *sql.DB, fsys fs.FS, args []string, out io.Writer) error {
	cmd, err := parseArgs(args)
	if err != nil {
		return err
	}

	src, err := iofs.New(fsys, ".")
	if err != nil {
		return fmt.Errorf("read migrations: %w", err)
	}
	driver, err := pgx.WithInstance(db, &pgx.Config{MigrationsTable: MigrationsTable})
	if err != nil {
		return fmt.Errorf("init migration driver: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", src, "pgx5", driver)
	if err != nil {
		return fmt.Errorf("init migrator: %w", err)
	}
	defer m.Close()

	switch cmd.name {
	case "up":
		err = m.Up()
	case "down":
		if cmd.all {
			err = m.Down()
		} else {
			err = m.Steps(-cmd.steps)
		}
	case "force":
		err = m.Force(cmd.force)
	}
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate %s: %w", cmd.name, err)
	}

	version, dirty, err := m.Version()
	switch {
	case errors.Is(err, migrate.ErrNilVersion):
		fmt.Fprintln(out, "schema version: none (no migrations applied)")
	case err != nil:
		return fmt.Errorf("read schema version: %w", err)
	default:
		fmt.Fprintf(out, "schema version: %d (dirty: %t)\n", version, dirty)
	}
	return nil
}
