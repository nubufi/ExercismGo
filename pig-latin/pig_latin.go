package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	var pigLatinWords []string
	for _, word := range words {
		pigLatinWords = append(pigLatinWords, translateWord(word))
	}
	return strings.Join(pigLatinWords, " ")
}

func translateWord(word string) string {
	if translated, ok := rule1(word); ok {
		return translated
	}
	if translated, ok := rule3(word); ok {
		return translated
	}
	if translated, ok := rule4(word); ok {
		return translated
	}
	if translated, ok := rule2(word); ok {
		return translated
	}
	return word
}

func rule1(word string) (string, bool) {
	// If a word begins with a vowel, or starts with "xr" or "yt", add an "ay" sound to the end of the word.
	pattern := `^(?:[aeiouAEIOU]|xr|yt)`
	re := regexp.MustCompile(pattern)
	if re.MatchString(word) {
		return word + "ay", true
	} else {
		return word, false
	}
}

func rule2(word string) (string, bool) {
	// If a word begins with a one or more consonants, first move those consonants to the end of the word and then add an "ay" sound to the end of the word.
	consonantPattern := `^[^aeiouAEIOU]+`

	re := regexp.MustCompile(consonantPattern)

    // Find the leading consonants
    match := re.FindString(word)
	if match == "" {
		return word, false
	} else {
		return word[len(match):] + match + "ay", true
	}
}

func rule3(word string) (string, bool) {
	// If a word starts with zero or more consonants followed by "qu", first move those consonants (if any) and the "qu" part to the end of the word, and then add an "ay" sound to the end of the word.

	pattern := `^([^aeiouAEIOU]*qu)`

	re := regexp.MustCompile(pattern)

	// Find the leading consonants
	match := re.FindString(word)
	if match == "" {
		return word, false
	} else {
		return word[len(match):] + match + "ay", true
	}

}

func rule4(word string) (string, bool) {
	// If a word starts with one or more consonants followed by "y", first move the consonants preceding the "y"to the end of the word, and then add an "ay" sound to the end of the word.

	pattern := `^([^aeiouAEIOU]+)y`

	re := regexp.MustCompile(pattern)

	// Find the leading consonants
	match := re.FindString(word)
	if match == "" {
		return word,false
	} else {
		remaining := word[len(match)-1:]
		return remaining + match[:len(match)-1] + "ay", true
	}
}
