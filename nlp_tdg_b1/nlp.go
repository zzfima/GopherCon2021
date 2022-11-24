package gopherCon2021

import (
	"regexp"
	"strings"
)

var (
	wordRe   = regexp.MustCompile("[[:alpha:]]+")
	suffixes = []string{"s", "ing"}
)

// Stem return stemmed word
func Stem(word string) string {
	for _, s := range suffixes {
		if strings.HasSuffix(word, s) {
			return word[:len(word)-len(s)]
		}
	}

	return word
}

// Tokenize returns a slice of tokens found in text
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	// var tokens []string
	tokens := make([]string, 0, 20)
	for _, w := range words {
		// TODO: stem
		token := Stem(strings.ToLower(w))
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
