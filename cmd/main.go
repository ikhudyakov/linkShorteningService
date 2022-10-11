package main

import (
	"linkShorteningService/internal/database"
	"linkShorteningService/internal/handlers"
	u "linkShorteningService/internal/utility"

	_ "github.com/lib/pq"
)

func main() {
	err := database.Connect()
	if err != nil {
		u.CheckError(err)
		return
	}

	handlers.HandlersInit()
}
