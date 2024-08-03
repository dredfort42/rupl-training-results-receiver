package db

import (
	"database/sql"

	cfg "github.com/dredfort42/tools/configreader"
	_ "github.com/lib/pq"
)

// Database is the database struct
type Database struct {
	database             *sql.DB
	tableTrainingResults string
}

var db Database

// DatabaseInit initializes the database
func DatabaseInit() {
	db.tableTrainingResults = cfg.Config["db.table.training.results"]
	if db.tableTrainingResults == "" {
		panic("db.table.training.results is empty")
	}

	databaseConnect()
	tablesCheck()
}
