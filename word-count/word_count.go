package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}
	phrase = strings.ToLower(phrase)
	phrase = strings.ReplaceAll(phrase, ",", " ")

	for _, s := range strings.Fields(phrase) {
		word := extractWord(s)
		freq[word]++
	}

	return freq
}

func extractWord(s string) string {
	// Remove non-alphanumeric characters except apostrophes
	re := regexp.MustCompile("[^a-zA-Z0-9']+")
	cleanStr := re.ReplaceAllString(s, "")

	// Remove apostrophes at the beginning or end
	cleanStr = strings.Trim(cleanStr, "'")

	return cleanStr
}
