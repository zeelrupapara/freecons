package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Database Connection
func Database() (*sql.DB, error) {
	// Database connection
	db, err := sql.Open("sqlite3", os.ExpandEnv("${DB_NAME}"))
	if err != nil {
		log.Fatal("Database Not Connected Due To: ", err)
		return nil, err
	}
	maxConnetions, err := strconv.Atoi(os.ExpandEnv("${MAX_DB_OPEN_CONNECTIONS}"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.SetMaxOpenConns(maxConnetions)
	return db, nil
}
