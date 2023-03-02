package sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"errors"
)

type Database struct {
	db *sql.DB
}

func Setup() (*Database, error) {
	db, err := sql.Open("sqlite3", "./urls.db")
	if err != nil {
		return nil, err
	}

	// Create the table if it doesn't exist
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, short_url TEXT UNIQUE, long_url TEXT)")
    if err != nil {
        return nil, err
    }

	return &Database{db: db}, nil
}

func (d *Database) QueryLongURL(shortURL string) (string, error) {
	var longURL string
	err := db.QueryRow("SELECT long_url FROM urls WHERE short_url = ?", shortURL).Scan(&longURL)
	if err == sql.ErrNoRows {
	    return "", errors.New("Shortened URL does not exist")
	} else if err != nil {
	    return "", err
	}

	return longURL, nil
}

func (d *Database) QueryShortURL(longURL string) (string, error) {
	var shortURL string
	err := db.QueryRow("SELECT short_url FROM urls WHERE long_url = ?", longURL).Scan(&shortURL)
	if err == sql.ErrNoRows {
		return "", errors.New("Long URL does not exist")
	} else if err != nil {
	    return "", err
	}

	return shortURL, nil
}

func (d *Database) InsertShortenedURL(id, shortURL string, longURL string) (error) {
    if _, err := db.Exec("INSERT INTO urls (id, short_url, long_url) VALUES (?, ?, ?)", id, shortURL, longURL); err != nil {
        return err
    }

    return nil
}