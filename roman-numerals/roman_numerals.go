package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("out of range")
	}
	letters := []string{"M", "CM", "D", "CD","C", "XC", "L", "XL","X", "IX", "V", "IV","I"}
	lookUp := []int{1000, 900, 500, 400,100, 90, 50, 40,10, 9, 5, 4,1}

	roman := ""

	for i := 0; i < len(lookUp); i++ {
		for input >= lookUp[i] {
			roman += letters[i]
			input -= lookUp[i]
		}
	}

	return roman, nil
}
