package handlers

import (
	"encoding/json"
	"fmt"
	db "linkShorteningService/internal/database"
	"linkShorteningService/internal/repo"
	u "linkShorteningService/internal/utility"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func CreateShortLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var link repo.Link = *repo.GetLink()
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		u.CheckError(err)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = url.ParseRequestURI(link.FullLink)
	if err != nil {
		u.CheckError(err)
		w.Write([]byte(err.Error()))
		return
	}

	shortLink, domain, err := db.GetShortLink(link.FullLink, link.Domain)
	if err != nil {
		u.CheckError(err)
		w.Write([]byte(err.Error()))
		return
	}

	if shortLink != "" {
		link.ShortLink = fmt.Sprintf("%s%s/%s", domain, repo.GetPort(), shortLink)
	} else {
		for {
			shortLink = link.Generate()
			check, err := db.CheckShortLink(shortLink)
			if err != nil {
				u.CheckError(err)
				w.Write([]byte(err.Error()))
				return
			}
			if !check {
				break
			}
			log.Println(shortLink)
		}

		link.ShortLink = shortLink
		lastId, lastDomain, err := db.SetLink(link)
		if err != nil {
			u.CheckError(err)
			w.Write([]byte(err.Error()))
			return
		}
		log.Println("set db with id =", lastId)
		link.ShortLink = fmt.Sprintf("%s%s/%s", lastDomain, repo.GetPort(), shortLink)
	}

	json.NewEncoder(w).Encode(link)
}

func GetFullLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	shortLink := params["shortlink"]
	link, err := db.GetFullLink(shortLink)
	if err != nil {
		u.CheckError(err)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, link, http.StatusSeeOther)
}

func HandlersInit() {
	r := mux.NewRouter()
	r.HandleFunc("/", CreateShortLink).Methods("POST")
	r.HandleFunc("/{shortlink}", GetFullLink).Methods("GET")

	log.Fatal(http.ListenAndServe(repo.GetPort(), limit(r)))
}
