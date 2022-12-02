package translation

import "strings"

func Translate(word string, language string) string {
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
	default:
		return ""
	}
}

func sanitizeInput(in string) string {
	out := strings.ToLower(in)
	return strings.TrimSpace(out)
}
