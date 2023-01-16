// Package translation contains the core service
package translation

import "strings"

type StaticService struct{}

func NewStaticService() *StaticService {
	return &StaticService{}
}

func (s *StaticService) Translate(word string, language string) string {
	// sanitize input
	word = sanitizeInput(word)
	language = sanitizeInput(language)

	if word != "hello" {
		return ""
	}
	switch language {
	case "english":
		return "hello"
	case "german":
		return "hallo"
	case "finnish":
		return "hei"
	case "french":
		return "bonjour"
	case "latin":
		return "ave"
	default:
		return ""
	}
}

func sanitizeInput(in string) string {
	out := strings.ToLower(in)
	return strings.TrimSpace(out)
}
