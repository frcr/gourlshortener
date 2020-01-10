package dbutil

import "github.com/frcr/gourlshortener/textutil"

// SaveURL creates the identifier and saves the URL into the database and returns
// its identifier
func SaveURL(url string, db DBQueryExecutable) string {
	i := 3
	var err error
	for i <= 10 {
		identifier := textutil.GenerateRandString(i)
		err = StoreIntoDB(identifier, url, db)
		if err != nil {
			i++ // TODO: Replace with more sensible collision resolution system
		} else {
			return identifier
		}
	}
	panic(err)

}
