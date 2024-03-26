package Database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var database *sql.DB

// GetInstance returns the singleton instance of the database connection.
// If the instance doesn't exist, it creates a new one and initializes the necessary tables.
func GetInstance() *sql.DB {
	if database == nil {
		once.Do(func() {
			var err error
			database, err = sql.Open("postgres", getConnectionString())
			if err != nil {
				panic("couldn't open database")
			}

			database.SetMaxOpenConns(250)

			initUsersTable(database)
		})
	}
	return database
}

// getConnectionString returns the connection string for the PostgreSQL database.
func getConnectionString() string {
	var (
		host     = os.Getenv("POSTGRES_HOSTNAME")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

// initUsersTable initializes the users table in the database if it doesn't already exist.
// It takes a *sql.DB as a parameter and executes the necessary SQL statements to create the table.
// If an error occurs during the execution, it panics.
func initUsersTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id			SERIAL PRIMARY KEY,
		email		varchar(32) UNIQUE,
		username	varchar(32) UNIQUE,
		password	varchar(84)
	);`)

	if err != nil {
		panic(err)
	}
}
