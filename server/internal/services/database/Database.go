package Database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var database *sql.DB

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

// takes an error from the postgres database
// returns a nicely formatted error
func ParseError(err error) error {
	msg := strings.TrimPrefix(err.Error(), "pq:")
	msg = strings.TrimSpace(msg)

	if strings.Contains(msg, "duplicate key") {
		parts := strings.Split(msg, "\"")
		if len(parts) < 2 {
			return err
		}
		dbfield := parts[1]

		fieldSlice := strings.Split(dbfield, "_")
		if len(fieldSlice) < 2 {
			return err
		}

		field := fieldSlice[1]

		msg = fmt.Sprintf("%s already in use", field)
	}

	return errors.New(msg)
}
