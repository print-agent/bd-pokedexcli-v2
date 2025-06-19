package pkg

import "strings"

func cleanInput(iText string) []string {
	iWords := strings.Fields(iText)
	return iWords
}
