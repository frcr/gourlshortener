package dbutil

// StoreIntoDB attempts to store the given id and url into DB
// and returns the error if any
func StoreIntoDB(identifier, url string, db DBQueryExecutable) error {
	_, err := db.Exec(insertValuesQuery, identifier, url)
	return err
}

func retrieveSingle(parameter, query string, db DBQueryExecutable) (string, error) {
	row := db.QueryRow(query, parameter) // QueryRow handles sanitizing the parameter
	var result string
	err := row.Scan(&result)
	return result, err
}

// RetrieveByID returns the URL corresponding to the identifier
// or an error if none was found
func RetrieveByID(identifier string, db DBQueryExecutable) (string, error) {
	return retrieveSingle(identifier, retrieveByIDQuery, db)
}

// RetrieveByURL returns the identifier corresponding to the URL
// or an error if none was found
func RetrieveByURL(url string, db DBQueryExecutable) (string, error) {
	return retrieveSingle(url, retrieveByURLQuery, db)
}
