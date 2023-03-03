package server

import (
	"strconv"

	sqlDB "UrlShortner/database"
	"UrlShortner/common"
)

type Server struct {
	database sqlDB.Databases
}

func Setup(db *sqlDB.Database) (*Server, error) {
	return &Server{database: db}, nil
}

func (s *Server) Redirect(shortURL string) (string, error){
    // Query the database for the long URL
    var longURL string
    
    longURL, err := s.database.QueryLongURL(shortURL)
    if err != nil {
    	return "", err
    }

    return longURL, nil
}

func (s *Server) Shorten(longURL string) (string, error) {

    var shortURL string
    shortURL, err := s.database.QueryShortURL(longURL)
    if err == nil {
        return shortURL, nil
    }

    // Generate a random short URL
    var id int64
    id = common.GenerateRandomNumber()
    shortURL = strconv.FormatInt(id, 36)
    if s.database.InsertShortenedURL(id, shortURL, longURL); err != nil {
    	return "", err
    }

    return shortURL, nil
}