// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimRightFunc(remark,unicode.IsSpace)

	silence := isSilence(remark)
	question := isQuestion(remark)
	yell := isYell(remark)
	
	if silence {
		return "Fine. Be that way!"
	} else if question && yell{
		return "Calm down, I know what I'm doing!" 
	} else if question {
		return "Sure."
	} else if yell {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

func isSilence(remark string) bool {
	for _,r := range remark {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func isQuestion(remark string) bool {
	if len(remark) == 0 {
		return false
	}

	return remark[len(remark)-1] == '?'
}

func isYell(remark string) bool {
	for _,r := range remark {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}

	return true && hasLetter(remark)
}

func hasLetter(remark string) bool {
	for _,r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false

}
