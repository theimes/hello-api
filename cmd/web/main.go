package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/theimes/hello-api/handlers"
	"github.com/theimes/hello-api/handlers/rest"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":4000"
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s", addr)

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
