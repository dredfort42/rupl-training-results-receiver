package db

import (
	"database/sql"

	cfg "github.com/dredfort42/tools/configreader"
	_ "github.com/lib/pq"
)

// Database is the database struct
type Database struct {
	Database              *sql.DB
	TableTrainingSessions string
}

var DB Database

// DatabaseInit initializes the database
func DatabaseInit() {
	DB.TableTrainingSessions = cfg.Config["db.table.training.sessions"]
	if DB.TableTrainingSessions == "" {
		panic("db.table.training.sessions is empty")
	}

	databaseConnect()
	tablesCheck()
}
