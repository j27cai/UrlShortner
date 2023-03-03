package server

import (
    "errors"
    "testing"
)

type mockDatabase struct{}

func (m *mockDatabase) QueryLongURL(shortURL string) (string, error) {
    if shortURL != "invalid" {
        return "long", nil
    }
    return "", errors.New("invalid")
}

func (m *mockDatabase) QueryShortURL(longURL string) (string, error) {
    if longURL != "invalid" {
        return "short", nil
    }
    return "", errors.New("invalid")
}

func (m *mockDatabase) InsertShortenedURL(id int64, shortURL string, longURL string) error {
    return nil
}

func TestRedirect(t *testing.T) {
    s := &Server{database: &mockDatabase{}}

    expectedLongURL := "long"

    longURL, err := s.Redirect("short")
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if longURL != expectedLongURL {
        t.Errorf("expected long")
    }

    _, err = s.Redirect("invalid")
    if err == nil {
        t.Error("expected error")
    }
}


func TestShorten(t *testing.T) {
    s := &Server{database: &mockDatabase{}}

    expectedShortURL := "long"

    shortURL, err := s.Redirect("long")
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    if shortURL != expectedShortURL {
        t.Errorf("expected short")
    }

    _, err = s.Redirect("invalid")
    if err == nil {
        t.Error("expected error")
    }
}
