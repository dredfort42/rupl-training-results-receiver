package db

import (
	"database/sql"

	cfg "github.com/dredfort42/tools/configreader"
	_ "github.com/lib/pq"
)

// Database is the database struct
type Database struct {
	Database             *sql.DB
	TableTrainingResults string
}

var DB Database

// DatabaseInit initializes the database
func DatabaseInit() {
	DB.TableTrainingResults = cfg.Config["db.table.training.results"]
	if DB.TableTrainingResults == "" {
		panic("db.table.training.results is empty")
	}

	databaseConnect()
	tablesCheck()
}
