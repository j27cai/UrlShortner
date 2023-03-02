package main
import (
    "database/sql"
    "fmt"
    "log"
    "math/rand"
    "net/http"

    "client"
    "common"
    "sql"
    "server"

    _ "github.com/mattn/go-sqlite3"
)

type url struct {
	OriginalURL string `json:"url"`
	ShortURL    string `json:"shortUrl"`
}

var db *sql.DB

func main() {
    // Open the database
    var err error
    db, err = sql.Open("sqlite3", "./urls.db")
    if err != nil {
        log.Fatal(err)
    }

    // Create the table if it doesn't exist
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, short_url TEXT UNIQUE, long_url TEXT)")
    if err != nil {
        log.Fatal(err)
    }

    // Handle requests
    http.HandleFunc("/", redirectHandler)
    http.HandleFunc("/shorten", shortenHandler)
    http.HandleFunc("/list", listHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
// For local development only
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
