package sql

import (
    "database/sql"
    "testing"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func SetupTest() error {
    var err error
    db, err = sql.Open("sqlite3", "./urls.db")
    if err != nil {
        return err
    }

    _, err = db.Exec("DROP TABLE IF EXISTS urls")
	if err != nil {
		return err
	}

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, short_url TEXT UNIQUE, long_url TEXT)")
    if err != nil {
        return err
    }

    return nil
}

func TestDatabase_QueryLongURL(t *testing.T) {
    err := SetupTest()
    if err != nil {
        t.Fatalf("failed to setup test: %v", err)
    }
    defer db.Close()

    db := &Database{db: db}
    shortURL := "abc123"
    expectedLongURL := "https://www.google.com"
    err = db.InsertShortenedURL("1", shortURL, expectedLongURL)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    longURL, err := db.QueryLongURL(shortURL)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if longURL != expectedLongURL {
        t.Errorf("unexpected long URL: got %q, want %q", longURL, expectedLongURL)
    }
}

func TestDatabase_QueryShortURL(t *testing.T) {
    err := SetupTest()
    if err != nil {
        t.Fatalf("failed to setup test: %v", err)
    }
    defer db.Close()

    db := &Database{db: db}
    shortURL := "abc123123"
    longURL := "https://www.google.coms"
    err = db.InsertShortenedURL("1", shortURL, longURL)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    foundShortURL, err := db.QueryShortURL(longURL)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if foundShortURL != shortURL {
        t.Errorf("unexpected short URL: got %q, want %q", foundShortURL, shortURL)
    }
}

func TestDatabase_InsertShortenedURL(t *testing.T) {
    err := SetupTest()
    if err != nil {
        t.Fatalf("failed to setup test: %v", err)
    }
    defer db.Close()

    db := &Database{db: db}
    // Insert a new shortened URL
    err = db.InsertShortenedURL("123", "short", "long")
    if err != nil {
        t.Fatal(err)
    }

    // Query the shortened URL to verify that it was inserted correctly
    shortURL, err := db.QueryShortURL("long")
    if err != nil {
        t.Fatal(err)
    }

    // Check that the retrieved shortened URL matches the expected value
    expectedShortURL := "short"
    if shortURL != expectedShortURL {
        t.Errorf("QueryShortURL returned %s, expected %s", shortURL, expectedShortURL)
    }

    // Query the long URL to verify that it was inserted correctly
    longURL, err := db.QueryLongURL("short")
    if err != nil {
        t.Fatal(err)
    }

    // Check that the retrieved long URL matches the expected value
    expectedLongURL := "long"
    if longURL != expectedLongURL {
        t.Errorf("QueryLongURL returned %s, expected %s", longURL, expectedLongURL)
    }
}
