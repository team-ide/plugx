package databasex

import (
	"database/sql"
)

const (
	PluginName = "database"
)

type Plugin interface {
	Open(c string) (db *sql.DB, err error)
	IsDatabasePlugin()
}
