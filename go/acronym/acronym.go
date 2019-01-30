// Package acronym abbreviates phrases
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate turns a phrase into its acronym
func Abbreviate(s string) string {
	re := regexp.MustCompile(`[\s\-]+`)
	words := re.Split(s, -1)
	result := ""
	for _, word := range words {
		result += strings.ToUpper(word[0:1])
	}
	return result
}
