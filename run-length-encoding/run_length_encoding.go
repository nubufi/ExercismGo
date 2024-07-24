package encode

import (
	"fmt"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	if len(input) <=1 {
		return input
	}
	result := ""
	n := 0
	for n < len(input) {
		c := input[n]
		count := 1
		i := n + 1
		if i < len(input) {
			for c == input[i] {
				count++
				i++
				if i == len(input) {break}
			}
		}

		if count > 1 {
			result += fmt.Sprintf("%d%c", count, c)
		} else {
			result += string(c)
		}

		n += count
	}

	return result
}

func RunLengthDecode(input string) string {
	var result string

	if len(input) <= 1 {
		return input
	}
	
	countStr := ""
	for i := range input {
		c := input[i]
		if !isNumber(c) {
			if countStr == "" {
				result += string(c)
			} else {
				count,_ := strconv.Atoi(countStr)
				result += strings.Repeat(string(c),count)
			}
			countStr = ""

		} else {
			countStr += string(c)
		}
	}
	return result
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}
