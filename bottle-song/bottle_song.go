package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	numMap := []string{"No", "One", "Two", "Three", "Four", "Five",	"Six", "Seven", "Eight", "Nine", "Ten"}

	var lyrics []string
	var remaining,l1,l4 string

	l3 := "And if one green bottle should accidentally fall,"
	for i:=0;i<takeDown;i++ {
		if startBottles == 1 {
			l1 = fmt.Sprintf("%s green bottle hanging on the wall,",numMap[startBottles])
		}else {
			l1 = fmt.Sprintf("%s green bottles hanging on the wall,",numMap[startBottles])
		}
		if startBottles-1 == 0 {
			remaining = "no"
		} else {
			remaining = strings.ToLower(numMap[startBottles-1])
		}
		if startBottles-1 != 1 {
			l4 = fmt.Sprintf("There'll be %s green bottles hanging on the wall.",remaining)
		} else {
			l4 = fmt.Sprintf("There'll be %s green bottle hanging on the wall.",remaining)
		}
		lyrics = append(lyrics,l1,l1,l3,l4)
		if i != takeDown-1 {
			lyrics = append(lyrics,"")
		}
		startBottles--
	}

	return lyrics
}
