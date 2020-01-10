package dbutil

// CreateTableQuery tries to create table when connecting to DB
const createTableQuery string = "CREATE TABLE IF NOT EXISTS  urls (id VARCHAR(20) PRIMARY KEY, url TEXT)"

// InsertValuesQuery tries to insert values into the DB
const insertValuesQuery string = "INSERT INTO urls (id, url) VALUES ($1, $2)"

// RetrieveByIDQuery tries to find the URL corresponding to the id
const retrieveByIDQuery string = "SELECT url FROM urls WHERE id = $1"

// RetrieveByURLQuery tries to find the id corresponding to this url
const retrieveByURLQuery string = "SELECT id FROM urls WHERE url = $1"
