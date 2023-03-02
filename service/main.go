package main
import (
    "database/sql"
    // "fmt"
    "log"
    "net/http"

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

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://google.com", http.StatusFound)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Write([]byte("http://localhost:8080/" + "abcdef"))
}

// List all the current shortened URLs and Long Urls. Helpful endpoint so there's no need to keep track of everything
// Ideally, this endpoint would turn into an universal get handler for urls based on users, pagination, and order by
func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Write([]byte("http://localhost:8080/" + "abcdef"))
}
