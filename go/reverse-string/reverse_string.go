package reverse

import "strings"

// String reverses a string
func String(s string) string {
	cs := strings.Split(s, "")
	for left, right := 0, len(cs)-1; left < right; left, right = left+1, right-1 {
		cs[left], cs[right] = cs[right], cs[left]
	}
	return strings.Join(cs, "")
}
