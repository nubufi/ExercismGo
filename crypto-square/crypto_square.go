package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	pt = normalize(pt)
	if pt == "" {
		return ""
	}
	r,c := calcRC(len(pt))

	if r*c > len(pt) {
		pt += strings.Repeat(" ", r*c-len(pt))
	}

	splitted := splitToRows(pt, c)
	var res string
	for i := 0; i < c; i++ {
		for _, row := range splitted {
			res += string(row[i])
		}
		if i < c-1 {
			res += " "
		}
	}

	return res
}

func splitToRows(s string, c int) []string {
	var res []string
	for i := 0; i < len(s); i += c {
		end := i + c
		res = append(res, s[i:end])
	}
	return res
}

func calcRC(l int) (r, c int) {
	r = int(math.Sqrt(float64(l)))
	c = l / r

	if c-r>1 {
		r++
		c--
	}
	
	if r*c < l {
		if r < c {
			r++
		} else {
			c++
		}
	}

	return r, c
}

func normalize(s string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	return re.ReplaceAllString(strings.ToLower(s), "")
}
