package server

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"linkShorteningService/internal/config"
	"linkShorteningService/internal/repo"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestDBManager struct {
	DB *sql.DB
}

func (db *TestDBManager) GetShortLink(link string, domainId int) (string, string, error) {
	return "1q2dr3g4jJ", "localhost:8001", nil
}
func (db *TestDBManager) CheckShortLink(shortLlink string) (bool, error) {
	return false, nil
}
func (db *TestDBManager) SetLink(link repo.Link) (int64, string, error) {
	return 1, "localhost:8001", nil
}
func (db *TestDBManager) GetFullLink(shortLink string) (string, error) {
	return "https://test1.com", nil
}

func TestCreateShortLink(t *testing.T) {
	body := repo.Link{}
	body.FullLink = "https://test1.com"
	body.Domain = 0
	var dbmanager repo.DBmanager = &TestDBManager{}

	conf := config.Config{
		Host: "localhost",
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body)

	req, err := http.NewRequest("POST", "/", buffer)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateShortLink(w, r, dbmanager, &conf)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedLink := repo.Link{
		FullLink:  "https://test1.com",
		ShortLink: "localhost:8001/1q2dr3g4jJ",
		Domain:    0,
	}

	gotLink := repo.Link{}
	json.NewDecoder(rr.Body).Decode(&gotLink)

	assert.Equal(t, expectedLink, gotLink)
}

func TestGetFullLink(t *testing.T) {
	var dbmanager repo.DBmanager = &TestDBManager{}
	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("shortlink", "1q2dr3g4jJ")
	buffer.WriteString(params.Encode())
	req, err := http.NewRequest("GET", "/", buffer)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetFullLink(w, r, dbmanager)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
