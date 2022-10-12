package repo

import (
	"database/sql"
	db "linkShorteningService/internal/database"
	"log"
)

// Получение короткой ссылки из базы
func GetShortLink(link string, domainId int) (string, string, error) {
	var shortLink string
	var domain string
	var err error
	var rows *sql.Rows

	db := db.GetDB()

	rows, err = db.Query("select shortlink, d.domain from links l join domains d on l.domain=d.id where l.link = $1 and l.domain = $2", link, domainId)
	if err != nil {
		log.Println((err.Error()))
		return shortLink, domain, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&shortLink, &domain)
		if err != nil {
			log.Println((err.Error()))
			return shortLink, domain, err
		}
	}

	return shortLink, domain, err
}

// проверка, существует ли уже такая же ссылка в базе
func CheckShortLink(shortLlink string) (bool, error) {
	var err error
	var rows *sql.Rows
	var shortLlinkFromDB string

	db := db.GetDB()

	rows, err = db.Query("select shortlink from links")
	if err != nil {
		log.Println((err.Error()))
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&shortLlinkFromDB)
		if err != nil {
			log.Println((err.Error()))
			return false, err
		}
		if shortLlink == shortLlinkFromDB {
			return true, err
		}
	}

	return false, err
}

// Запись новой ссылки в базу
func SetLink(link Link) (int64, string, error) {
	var lastID int64
	var domain string
	var err error
	var rows *sql.Rows

	db := db.GetDB()

	err = db.QueryRow(
		"INSERT INTO links (link, shortlink, domain) VALUES ($1, $2, $3) RETURNING id",
		link.FullLink,
		link.ShortLink,
		link.Domain).Scan(&lastID)
	if err != nil {
		return lastID, domain, err
	}

	rows, err = db.Query("SELECT d.domain FROM domains d JOIN links l ON d.id=l.domain WHERE l.id=$1", lastID)
	if err != nil {
		return lastID, domain, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&domain)
		if err != nil {
			return lastID, domain, err
		}
	}

	return lastID, domain, err
}

// Получение полной ссылки из базы
func GetFullLink(shortLink string) (string, error) {
	var link string
	var err error
	var rows *sql.Rows

	db := db.GetDB()

	rows, err = db.Query("select link from links where shortlink = $1", shortLink)
	if err != nil {
		log.Println((err.Error()))
		return link, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&link)
		if err != nil {
			log.Println((err.Error()))
			return link, err
		}
	}
	return link, err
}
