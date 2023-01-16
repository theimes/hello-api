// Package rest contains the web handler for the services
package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

type Translator interface {
	Translate(word string, language string) string
}

type TranslateHandler struct {
	service Translator
}

func New(service Translator) *TranslateHandler {
	return &TranslateHandler{
		service: service,
	}
}

func (t *TranslateHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	words := strings.Split(r.URL.Path, "/")

	word := words[len(words)-1]
	translation := t.service.Translate(word, language)

	if translation == "" {
		//language = ""
		log.Print("got no translation")
		w.WriteHeader(404)
		return
	}

	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
