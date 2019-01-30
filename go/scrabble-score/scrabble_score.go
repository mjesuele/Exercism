// Package scrabble calculates Scrabble scores
package scrabble

import "strings"

func letterScore(l string) int {
	switch strings.ToUpper(l) {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	default:
		return 0
	}
}

// Score calculates the Scrabble score of a given word
func Score(s string) int {
	score := 0
	for _, l := range s {
		score += letterScore(string(l))
	}
	return score
}
