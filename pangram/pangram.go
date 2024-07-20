package pangram

import "unicode"

func IsPangram(input string) bool {
	letterSet := make(map[rune]bool)
	for _, letter := range input {
		lowerChar := unicode.ToLower(letter)

		if lowerChar >= 'a' && lowerChar <= 'z' {
			letterSet[lowerChar] = true
		}
	}

	return len(letterSet) == 26
}
