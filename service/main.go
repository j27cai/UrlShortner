package main
import (
	"fmt"
    "log"
    "net/http"
    "strings"

    "UrlShortner/server"
    sqlDB "UrlShortner/database"
    _ "github.com/mattn/go-sqlite3"
)

// type url struct {
//  OriginalURL string `json:"url"`
//  ShortURL    string `json:"shortUrl"`
// }

var db *sqlDB.Database
var service *server.Server

func main() {
    // Open the database
    var err error
    db, err = sqlDB.Setup()
    if err != nil {
        log.Fatal(err)
        return
    }

    service, err = server.Setup(db)
    if err != nil {
    	log.Fatal(err)
    	return
    }

    http.HandleFunc("/", redirectHandler)
    http.HandleFunc("/shorten", shortenHandler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    // Get the short URL
    shortURL := strings.TrimPrefix(r.URL.Path, "/")
    if shortURL == "" {
        http.NotFound(w, r)
        return
    }

    longURL, err := service.Redirect(shortURL)
    if err != nil {
    	http.Error(w, "Error redirecting", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, longURL, http.StatusFound)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    // Get the long URL
    longURL := r.FormValue("url")
    if longURL == "" {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    shortURL, err := service.Shorten(longURL)
    if err != nil {
    	http.Error(w, "Error shortening", http.StatusInternalServerError)
        return
    }

    w.Write([]byte(fmt.Sprintf("%s/%s", r.Host, shortURL)))
}
   
