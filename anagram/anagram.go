package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var anagrams []string
	
	for _,c := range candidates {
		if isAnagram(subject,c) {
			anagrams = append(anagrams,c)
		}
	}

	return anagrams
}


func isAnagram(subject, candidate string) bool {
	if len(subject) != len(candidate){
		return false
	}
	

	subject = strings.ToLower(subject)
	candidate = strings.ToLower(candidate)
	
	if subject == candidate {
		return false
	}

	rune1 := []rune(subject)
	rune2 := []rune(candidate)

	// Sort the slices of runes
	sort.Slice(rune1, func(i, j int) bool {
		return rune1[i] < rune1[j]
	})
	sort.Slice(rune2, func(i, j int) bool {
		return rune2[i] < rune2[j]
	})

	
	for i := range rune1 {
		if rune1[i] != rune2[i] {
			return false
		}
	}

	return true
}
