package main

import (
	"linkShorteningService/internal/database"
	"linkShorteningService/internal/handlers"
	u "linkShorteningService/internal/utility"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.CreateShortLink).Methods("POST")
	r.HandleFunc("/{shortlink}", handlers.GetLink).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", r))
}
