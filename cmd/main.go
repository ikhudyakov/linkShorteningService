package main

import (
	"linkShorteningService/internal/database"
	"linkShorteningService/internal/server"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Println((err.Error()))
		return
	}

	server.ServerInit()
}
