package database

import (
	"database/sql"
	"fmt"
	"linkShorteningService/internal/repo"
	u "linkShorteningService/internal/utility"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func startConnection() error {
	host := u.GetEnv("HOST")
	port, _ := strconv.Atoi(u.GetEnv("PORT"))
	user := u.GetEnv("USER")
	password := u.GetEnv("PASSWORD")
	dbname := u.GetEnv("DBNAME")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Connect() error {
	var err error
	if db == nil {
		err = startConnection()
	}
	return err
}

func GetShortLink(link string, domainId int) (string, string) {
	rows, err := db.Query("select shortlink, d.domain from links l join domains d on l.domain=d.id where l.link = $1 and l.domain = $2", link, domainId)
	u.CheckError(err)
	defer rows.Close()
	var shortLink string
	var domain string

	for rows.Next() {
		err := rows.Scan(&shortLink, &domain)
		u.CheckError(err)
	}

	return shortLink, domain
}

func SetLink(link repo.Link) (int64, string) {
	var lastID int64
	var domain string
	err = db.QueryRow(
		"INSERT INTO links (link, shortlink, domain) VALUES ($1, $2, $3) RETURNING id",
		link.FullLink,
		link.ShortLink,
		link.Domain).Scan(&lastID)
	u.CheckError(err)

	rows, err := db.Query("SELECT d.domain FROM domains d JOIN links l ON d.id=l.domain WHERE l.id=$1", lastID)
	u.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&domain)
		u.CheckError(err)
	}

	u.CheckError(err)
	return lastID, domain
}

func GetLink(shortLink string) string {
	var link string
	rows, err := db.Query("select link from links where shortlink = $1", shortLink)
	u.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&link)
		u.CheckError(err)
	}
	return link
}
