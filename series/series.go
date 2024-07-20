package series

import "strings"

func All(n int, s string) []string {
	substrings := []string{}
	for i := 0; i <= len(s)-n; i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return substrings
}

func UnsafeFirst(n int, s string) string {
	f, _ := First(n, s)
	return f
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		padded := s + strings.Repeat("0", n-len(s))
		return padded, false
	}
	return s[:n], true
}
