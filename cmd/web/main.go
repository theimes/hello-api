package main

import (
	"log"
	"net/http"

	"github.com/theimes/hello-api/handlers/rest"
)

func main() {
	addr := ":4000"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s", addr)

	log.Fatal(http.ListenAndServe(addr, mux))

}
