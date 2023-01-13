package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/theimes/hello-api/handlers"
	"github.com/theimes/hello-api/handlers/rest"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = "4000"
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux))

}
