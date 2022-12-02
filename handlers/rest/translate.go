package rest

import (
	"encoding/json"
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

	language := "english"
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := translation.Translate(word, language)

	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
