// Package faas wraps the translate handler into a function as a Service wrapper
package faas

import (
	"net/http"

	"github.com/theimes/hello-api/handlers/rest"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
