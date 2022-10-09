package handlers

import (
	"encoding/json"
	db "linkShorteningService/internal/database"
	"linkShorteningService/internal/repo"
	u "linkShorteningService/internal/utility"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateShortLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var link repo.Link = *repo.GetLink()
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		u.CheckError(err)
		return
	}

	if shortLink, domain := db.GetShortLink(link.FullLink, link.Domain); shortLink != "" {
		link.ShortLink = domain + ":8001/" + shortLink //	исправить
	} else {
		shortLink := link.Generate()
		link.ShortLink = shortLink
		lastId, lastDomain := db.SetLink(link)
		log.Println("set db with id =", lastId)
		link.ShortLink = lastDomain + ":8001/" + shortLink //	исправить
	}

	json.NewEncoder(w).Encode(link)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	shortLink := params["shortlink"]
	link := db.GetLink(shortLink)
	http.Redirect(w, r, link, http.StatusSeeOther)
}
