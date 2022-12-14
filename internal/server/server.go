package server

import (
	c "linkShorteningService/internal/config"
	"linkShorteningService/internal/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlersInit(dbmanager repo.DBmanager, conf *c.Config) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		CreateShortLink(w, r, dbmanager, conf)
	}).Methods("POST")
	r.HandleFunc("/{shortlink}", func(w http.ResponseWriter, r *http.Request) {
		GetFullLink(w, r, dbmanager)
	}).Methods("GET")

	return r
}
