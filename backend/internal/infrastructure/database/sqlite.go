package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS voucher_assignments (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	crew_name TEXT NOT NULL,
	crew_id TEXT NOT NULL,
	flight_number TEXT NOT NULL,
	flight_date TEXT NOT NULL,
	aircraft_type TEXT NOT NULL,
	seat1 TEXT NOT NULL,
	seat2 TEXT NOT NULL,
	seat3 TEXT NOT NULL,
	created_at TEXT NOT NULL,
	UNIQUE(flight_number, flight_date)
);`

// NewDatabaseConnection initializes and migrates the SQLite database.
func NewDatabaseConnection() (*sql.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "vouchers.db" // Default for local dev
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Run migration
	if _, err = db.Exec(createTableQuery); err != nil {
		return nil, err
	}

	log.Printf("Database initialized at %s", dbPath)
	return db, nil
}
