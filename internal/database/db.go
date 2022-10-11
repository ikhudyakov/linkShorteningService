package database

import (
	"database/sql"
	"fmt"
	"linkShorteningService/internal/repo"
	u "linkShorteningService/internal/utility"

	_ "github.com/lib/pq"
)

var db *sql.DB

func startConnection() error {
	var err error

	conf, err := repo.GetConfig()
	if err != nil {
		return err
	}

	host := conf.Host
	port := conf.Postgresqlport
	user := conf.User
	password := conf.Password
	dbname := conf.DBname

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

// Получение короткой ссылки из базы
func GetShortLink(link string, domainId int) (string, string, error) {
	var shortLink string
	var domain string
	var err error
	var rows *sql.Rows

	rows, err = db.Query("select shortlink, d.domain from links l join domains d on l.domain=d.id where l.link = $1 and l.domain = $2", link, domainId)
	if err != nil {
		u.CheckError(err)
		return shortLink, domain, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&shortLink, &domain)
		if err != nil {
			u.CheckError(err)
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

	rows, err = db.Query("select shortlink from links")
	if err != nil {
		u.CheckError(err)
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&shortLlinkFromDB)
		if err != nil {
			u.CheckError(err)
			return false, err
		}
		if shortLlink == shortLlinkFromDB {
			return true, err
		}
	}

	return false, err
}

// Запись новой ссылки в базу
func SetLink(link repo.Link) (int64, string, error) {
	var lastID int64
	var domain string
	var err error
	var rows *sql.Rows

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

	rows, err = db.Query("select link from links where shortlink = $1", shortLink)
	if err != nil {
		u.CheckError(err)
		return link, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&link)
		if err != nil {
			u.CheckError(err)
			return link, err
		}
	}
	return link, err
}
