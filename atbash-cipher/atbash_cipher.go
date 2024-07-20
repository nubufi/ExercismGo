package atbash

import "strings"

func Atbash(s string) string {
	var res string
	var count int

	s = strings.ToLower(s)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			res += string('z' - r + 'a')
			count++
		} else if r >= '0' && r <= '9' {
			res += string(r)
			count++
		}
		if count%5 == 0 && count != 0 && res[len(res)-1] != ' ' {
			res += " "
		}
	}
	res = strings.TrimRight(res, " ")
	return res
}
