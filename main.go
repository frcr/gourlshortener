package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/frcr/gourlshortener/dbutil"
	"github.com/frcr/gourlshortener/textutil"
)

type urlpayload struct {
	URL string
}

func handleGet(w http.ResponseWriter, r *http.Request, db dbutil.DBQueryExecutable) {
	identifier := r.URL.RequestURI()[1:]
	if len(identifier) == 0 {
		fmt.Fprintf(w, textutil.RenderResponse(r.Host, ""))
		return
	}
	url, err := dbutil.RetrieveByID(identifier, db)
	if err != nil {
		fmt.Fprintf(w, textutil.RenderResponse(r.Host, "Sorry, this URL was not found. Try shortening another one!"))
		return
	}
	http.Redirect(w, r, url, 302)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, db dbutil.DBQueryExecutable) {
	var u urlpayload
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Fprintf(w, `{"error": "Error! Unable to parse the payload."}`)
		return
	}
	_, err = url.ParseRequestURI(u.URL)
	if err != nil {
		fmt.Fprintf(w, `{"error": "Error! This URL is not valid."}`)
		return
	}
	identifier, err := dbutil.RetrieveByURL(u.URL, db)
	if err != nil {
		identifier = dbutil.SaveURL(u.URL, db)
	}
	fmt.Fprintf(w, `{"success": "`+fmt.Sprintf("http://%s/%s", r.Host, identifier)+`"}`)
}

// POST "/"             -> generate shortened URL and return json with success info or
//                         return json with error
//
// GET  "/"             -> show index if empty or
//                         redirect to url by hash or
//                         show index with error
func primaryHandler(w http.ResponseWriter, r *http.Request, db dbutil.DBQueryExecutable) {
	switch r.Method {
	case "GET":
		handleGet(w, r, db)
	case "POST":
		handlePost(w, r, db)
	default:
		fmt.Fprintf(w,
			textutil.RenderResponse(r.Host, "Error! Only GET and POST methods are allowed for this endpoint."))
	}
}

func main() {
	db := dbutil.GetConnection()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		primaryHandler(w, r, db)
	})
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
