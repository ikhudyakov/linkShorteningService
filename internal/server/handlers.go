package server

import (
	"encoding/json"
	"fmt"
	c "linkShorteningService/internal/config"
	"linkShorteningService/internal/repo"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func CreateShortLink(w http.ResponseWriter, r *http.Request, db repo.DBmanager, conf *c.Config) {
	w.Header().Set("Content-Type", "application/json")

	var link repo.Link = *repo.GetLink()

	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		log.Println((err.Error()))
		w.Write([]byte(err.Error()))
		return
	}

	if _, err := url.ParseRequestURI(link.FullLink); err != nil {
		log.Println((err.Error()))
		w.Write([]byte(err.Error()))
		return
	}

	shortLink, domain, err := db.GetShortLink(link.FullLink, link.Domain)
	if err != nil {
		log.Println((err.Error()))
		w.Write([]byte(err.Error()))
		return
	}

	if shortLink != "" {
		link.ShortLink = fmt.Sprintf("%s%s/%s", domain, conf.Port, shortLink)
	} else {
		for {
			shortLink = link.Generate()
			check, err := db.CheckShortLink(shortLink)
			if err != nil {
				log.Println((err.Error()))
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
			log.Println((err.Error()))
			w.Write([]byte(err.Error()))
			return
		}
		log.Println("set db with id =", lastId)
		link.ShortLink = fmt.Sprintf("%s%s/%s", lastDomain, conf.Port, shortLink)
	}

	json.NewEncoder(w).Encode(link)
}

func GetFullLink(w http.ResponseWriter, r *http.Request, db repo.DBmanager) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	shortLink := params["shortlink"]
	link, err := db.GetFullLink(shortLink)
	if err != nil {
		log.Println((err.Error()))
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, link, http.StatusSeeOther)
}
