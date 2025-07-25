package pkg

import (
	"regexp"
	"strings"
)

func CleanInput(text string) []string {
	cleanText := strings.TrimSpace(text)
	if cleanText == "" {
		return []string{}
	}

	words := strings.Fields(cleanText)

	var cleanWords []string
	punctuationRegex := regexp.MustCompile(`[^\p{L}']+`)
	for _, word := range words {
		cleanWord := punctuationRegex.ReplaceAllString(word, "")
		if cleanWord != "" {
			cleanWord = strings.ToLower(cleanWord)
			cleanWords = append(cleanWords, cleanWord)
		}
	}
	return cleanWords
}
