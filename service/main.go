package main
import (
    "log"

    sqlDB "UrlShortner/database"
    _ "github.com/mattn/go-sqlite3"
)

// type url struct {
//  OriginalURL string `json:"url"`
//  ShortURL    string `json:"shortUrl"`
// }

var db *sqlDB.Database

func main() {
    // Open the database
    var err error
    db, err = sqlDB.Setup()
    if err != nil {
        log.Fatal(err)
    }
}
