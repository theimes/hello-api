package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theimes/hello-api/handlers/rest"
	"github.com/theimes/hello-api/translation"
)

func TestTranslate(t *testing.T) {

	// given a set of test cases
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
	// and a web handler
	translator := translation.NewStaticService()
	translateHandler := rest.New(translator)
	handler := http.HandlerFunc(translateHandler.TranslateHandler)

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
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		if err != nil {
			t.Logf("couldn't unmarshal response %s: %s", rr.Body.String(), err.Error())
		}

		if resp.Language != tt.ExpectedLanguage {
			t.Errorf(`expected language %q but got %q`, tt.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != tt.ExpectedTranslation {
			t.Errorf(`expected %q but got %q`, tt.ExpectedTranslation, resp.Translation)
		}
	}
}
