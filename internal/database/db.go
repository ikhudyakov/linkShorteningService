package database

import (
	"database/sql"
	"fmt"
	c "linkShorteningService/internal/repo/config"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect() error {
	var err error
	var conn string

	conf, err := c.GetConfig()
	if err != nil {
		return err
	}

	host := conf.Host
	port := conf.Postgresqlport
	user := conf.User
	password := conf.Password
	dbname := conf.DBname
	connType := conf.ConnectionType

	switch connType {
	case "postgres":
		conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		if db, err = sql.Open(connType, conn); err != nil {
			return err
		}
	case "sqlite3":
		conn = dbname
		if db, err = sql.Open(connType, conn); err != nil {
			return err
		}
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
