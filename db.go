package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func dbConnect() {
	host := getEnv("HOST")
	port, _ := strconv.Atoi(getEnv("PORT"))
	user := getEnv("USER")
	password := getEnv("PASSWORD")
	dbname := getEnv("DBNAME")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
	err = db.Ping()
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Println((err.Error()))
	}
}

func GetDB() *sql.DB {
	if db == nil {
		dbConnect()
	}
	return db
}

func GetShortLinkFromDB(link string) string {
	GetDB()
	rows, err := db.Query("select shortlink from links where link = $1", link)
	CheckError(err)
	defer rows.Close()
	var shortLink string

	for rows.Next() {
		err := rows.Scan(&shortLink)
		CheckError(err)
	}

	return shortLink
}

func SetLinkToDB(link Link) int64 {
	GetDB()
	var lastID int64
	err = db.QueryRow(
		"INSERT INTO links (link, shortlink) VALUES ($1, $2) RETURNING id",
		link.FullLink,
		link.ShortLink).Scan(&lastID)

	CheckError(err)
	return lastID
}

func GetLinkFromDB(shortLink string) string {
	GetDB()
	var link string
	rows, err := db.Query("select link from links where shortlink = $1", shortLink)
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&link)
		CheckError(err)
	}
	return link
}
