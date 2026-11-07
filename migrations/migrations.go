// Package migrations embeds the SQL migration files so the binary can apply
// and verify the schema without needing the migrations folder on disk.
package migrations

import "embed"

// FS holds every *.sql migration file (golang-migrate naming: NNN_name.up.sql / .down.sql).
//
//go:embed *.sql
var FS embed.FS
