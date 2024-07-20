// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	reg := regexp.MustCompile(`[^a-zA-Z']+`)
	s = reg.ReplaceAllString(s, " ")

	var abb string

	for _, c := range strings.Fields(s) {
		abb += string(c[0])
	}

	return abb
}
