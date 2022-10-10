package main

import (
	"linkShorteningService/internal/database"
	"linkShorteningService/internal/handlers"
	u "linkShorteningService/internal/utility"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	err := database.Connect()
	if err != nil {
		u.CheckError(err)
		return
	}

	handlers.HandlersInit()
}
