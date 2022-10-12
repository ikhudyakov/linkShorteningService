package main

import (
	"context"
	"database/sql"
	c "linkShorteningService/internal/config"
	"linkShorteningService/internal/database"
	"linkShorteningService/internal/repo"
	pg "linkShorteningService/internal/repo/postgresql"
	"linkShorteningService/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var dbmanager repo.DBmanager
	var db *sql.DB
	var err error

	if db, err = database.Connect(); err != nil {
		log.Println((err.Error()))
		return
	}

	dbmanager = &pg.PGmanager{
		DB: db,
	}

	conf, err := c.GetConfig()
	if err != nil {
		log.Println(err)
	}

	r := server.HandlersInit(dbmanager)
	server := &http.Server{
		Addr:    conf.Port,
		Handler: server.Limit(r),
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
