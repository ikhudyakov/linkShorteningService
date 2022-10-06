package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateShortLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var link Link
	_ = json.NewDecoder(r.Body).Decode(&link)

	if shortLink := GetShortLinkFromDB(link.FullLink); shortLink != "" {
		link.ShortLink = r.Host + "/" + shortLink
	} else {
		shortLink := Generate()
		link.ShortLink = shortLink
		log.Println("id =", SetLinkToDB(link))
		link.ShortLink = r.Host + "/" + shortLink
	}

	json.NewEncoder(w).Encode(link)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	shortLink := params["shortlink"]
	link := GetLinkFromDB(shortLink)
	http.Redirect(w, r, link, http.StatusSeeOther)
}
