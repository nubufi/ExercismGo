package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question = strings.TrimSuffix(question, "?")
	question = strings.ReplaceAll(question, "What is ", "")
	question = strings.ReplaceAll(question, "by", "")

	splitted := strings.Fields(question)
	a, _ := strconv.Atoi(splitted[0])

	if a == 0 {
		return 0, false
	}

	if len(splitted) == 1 {
		return a, true
	} else if len(splitted) == 3 {
		b, _ := strconv.Atoi(splitted[2])
		return operate(a, b, splitted[1])
	} else if len(splitted) == 5 {
		b, _ := strconv.Atoi(splitted[2])
		c, _ := strconv.Atoi(splitted[4])
		r, t := operate(a, b, splitted[1])
		if !t {
			return 0, false
		}
		return operate(r, c, splitted[3])
	}

	return 0, false
}

func operate(a, b int, op string) (int, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}
	switch op {
	case "plus":
		return a + b, true
	case "minus":
		return a - b, true
	case "multiplied":
		return a * b, true
	case "divided":
		return a / b, true
	default:
		return a, true
	}
}
