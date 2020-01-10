package dbutil

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	// Bind the postgres driver
	_ "github.com/lib/pq"
)

// DBQueryExecutable is a standardized DB interface
type DBQueryExecutable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// dbConfig is a type holding the necessary data to connect to a database
type dbConfig struct {
	Host     string
	Port     uint
	User     string
	Password string
	DBName   string
}

// GetConnection gets a new connection to the database
func GetConnection() DBQueryExecutable {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := dbConfig{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configuration.Host,
		configuration.Port,
		configuration.User,
		configuration.Password,
		configuration.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.Exec(createTableQuery)

	// defer db.Close()
	return db
}
