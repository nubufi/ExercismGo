package isbn

import "regexp"

func IsValidISBN(isbn string) bool {
	re := regexp.MustCompile(`-`)
	isbn = re.ReplaceAllString(isbn, "")

	if len(isbn) != 10 {
		return false
	}

	var sum int
	for i, c := range isbn {
		var n int
		if i != 9 && (c < '0' || c > '9') {
			return false
		}
		if i == 9 && c != 'X' && (c < '0' || c > '9') {
			return false
		}
		if c == 'X' {
			n = 10
		} else {
			n = int(c - '0')
		}
		
		sum += n * (10 - i)
	}

	return sum%11 == 0
}

