package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/theimes/hello-api/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	words := strings.Split(r.URL.Path, "/")

	word := words[len(words)-1]
	translation := translation.Translate(word, language)

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
