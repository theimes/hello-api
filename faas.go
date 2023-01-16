// Package faas wraps the translate handler into a function as a Service wrapper
package faas

import (
	"net/http"

	"github.com/theimes/hello-api/handlers/rest"
	"github.com/theimes/hello-api/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translateService := translation.NewStaticService()
	translateHandler := rest.New(translateService)
	translateHandler.TranslateHandler(w, r)
}
