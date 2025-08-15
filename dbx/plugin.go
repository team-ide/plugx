package dbx

import (
	"database/sql"
)

type Plugin interface {
	Open(c string) (db *sql.DB, err error)
	IsDbPlugin()
}
