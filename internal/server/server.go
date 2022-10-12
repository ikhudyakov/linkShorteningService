package server

import (
	"context"
	c "linkShorteningService/internal/repo/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func HandlersInit() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", CreateShortLink).Methods("POST")
	r.HandleFunc("/{shortlink}", GetFullLink).Methods("GET")

	return r
}

func ServerInit() {

	r := HandlersInit()
	server := &http.Server{
		Addr:    c.GetPort(),
		Handler: limit(r),
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
