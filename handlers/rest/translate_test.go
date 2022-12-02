package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theimes/hello-api/handlers/rest"
)

func TestTranslate(t *testing.T) {

	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          404,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	for _, tt := range tt {
		// Arrange
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, tt.Endpoint, nil)

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		if rr.Code != tt.StatusCode {
			t.Errorf("expected status %d but got %d", tt.StatusCode, rr.Code)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != tt.ExpectedLanguage {
			t.Errorf(`expected language %q but got %q`, tt.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != tt.ExpectedTranslation {
			t.Errorf(`expected %q but got %q`, tt.ExpectedTranslation, resp.Translation)
		}
	}
}
