package translation_test

import (
	"testing"

	"github.com/theimes/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	tt := []struct {
		word     string
		language string
		want     string
	}{
		{
			word:     "hello",
			language: "english",
			want:     "hello",
		},
		{
			word:     "hello",
			language: "german",
			want:     "hallo",
		},
		{
			word:     "hello",
			language: "German",
			want:     "hallo",
		},
		{
			word:     "Hello",
			language: "german",
			want:     "hallo",
		},
		{
			word:     "hello ",
			language: "german",
			want:     "hallo",
		},
		{
			word:     "bye",
			language: "german",
			want:     "",
		},
		{
			word:     "hello",
			language: "finnish",
			want:     "hei",
		},
		{
			word:     "hello",
			language: "dutch",
			want:     "",
		},
		{
			word:     "bye",
			language: "dutch",
			want:     "",
		},
	}

	for _, testCase := range tt {
		got := translation.Translate(testCase.word, testCase.language)
		if got != testCase.want {
			t.Errorf(`translate %q to %s: expected %q, got %q`, testCase.word, testCase.language, testCase.want, got)
		}
	}

}
