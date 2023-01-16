package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theimes/hello-api/handlers/rest"
)

type stubbedService struct{}

func (s *stubbedService) Translate(word string, language string) string {
	if word == "foo" {
		return "bar"
	}
	return ""
}

func TestTranslateAPI(t *testing.T) {

	// given a set of test cases
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/translate/foo",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "bar",
		},
		{
			Endpoint:            "/translate/foo?language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "bar",
		},
		{
			Endpoint:            "/baz",
			StatusCode:          404,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}
	// and a web handler
	translator := stubbedService{}
	translateHandler := rest.New(&translator)
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
		_ = json.Unmarshal(rr.Body.Bytes(), &resp)
		//if err != nil {
		//	t.Logf("couldn't unmarshal response %s: %s", rr.Body.String(), err.Error())
		//}

		if resp.Language != tt.ExpectedLanguage {
			t.Errorf(`expected language %q but got %q`, tt.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != tt.ExpectedTranslation {
			t.Errorf(`expected %q but got %q`, tt.ExpectedTranslation, resp.Translation)
		}
	}
}
