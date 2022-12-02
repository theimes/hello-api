package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theimes/hello-api/handlers/rest"
)

func TestTranslate(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	handler := http.HandlerFunc(rest.TranslateHandler)

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 but got %d", rr.Code)
	}

	var resp rest.Resp
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.Language != "english" {
		t.Errorf(`expected language "english" but got %q`, resp.Language)
	}

	if resp.Translation != "hello" {
		t.Errorf(`expected "hello" but got %q`, resp.Translation)
	}
}
