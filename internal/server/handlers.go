package server

import (
	"encoding/json"
	"fmt"
	"linkShorteningService/internal/repo"
	c "linkShorteningService/internal/repo/config"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func CreateShortLink(w http.ResponseWriter, r *http.Request) {
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

	shortLink, domain, err := repo.GetShortLink(link.FullLink, link.Domain)
	if err != nil {
		log.Println((err.Error()))
		w.Write([]byte(err.Error()))
		return
	}

	if shortLink != "" {
		link.ShortLink = fmt.Sprintf("%s%s/%s", domain, c.GetPort(), shortLink)
	} else {
		for {
			shortLink = link.Generate()
			check, err := repo.CheckShortLink(shortLink)
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
		lastId, lastDomain, err := repo.SetLink(link)
		if err != nil {
			log.Println((err.Error()))
			w.Write([]byte(err.Error()))
			return
		}
		log.Println("set db with id =", lastId)
		link.ShortLink = fmt.Sprintf("%s%s/%s", lastDomain, c.GetPort(), shortLink)
	}

	json.NewEncoder(w).Encode(link)
}

func GetFullLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	shortLink := params["shortlink"]
	link, err := repo.GetFullLink(shortLink)
	if err != nil {
		log.Println((err.Error()))
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, link, http.StatusSeeOther)
}
