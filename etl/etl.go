package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newMap := make(map[string]int)

	for k, v := range in {
		for _, c := range v {
			newMap[strings.ToLower(c)] = k
		}
	}

	return newMap
}
