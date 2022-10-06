package main

import (
	"log"
	"net/http"
	"os"

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
	// db := GetDB()
	// db.Ping()

	r := mux.NewRouter()

	r.HandleFunc("/", CreateShortLink).Methods("POST")
	r.HandleFunc("/{shortlink}", GetLink).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", r))

}

func getEnv(text string) string {
	get, ok := os.LookupEnv(text)
	if !ok {
		log.Println("not found environment variable")
		return ""
	}
	return get
}
