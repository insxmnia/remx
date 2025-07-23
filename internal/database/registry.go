package database

import (
	"database/sql"

	tr "github.com/tursodatabase/libsql-client-go/libsql"
)

// Used to init a custom turso registry for both local and remote development without having to change the code
func init() {
	sql.Register("turso_remote", tr.Driver{})
}
