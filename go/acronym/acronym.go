// Package acronym exposes functions to turn strings into acronyms
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate turns a string into its acronym:
// Abbreviate("read the friendly manual") -> "RTFM"
func Abbreviate(s string) string {
	acronym := ""
	words := strings.Split(regexp.MustCompile("[^\\w^\\s^\\-^\\_]").ReplaceAllString(s, ""), " ")
	for _, word := range words {
		acronym += string(word[0])
	}
	return strings.ToUpper(acronym)
}
