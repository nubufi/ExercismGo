package yacht

import (
	"sort"
)

func Score(dice []int, category string) int {
	var score int

	switch category {
	case "little straight":
		copySlice := make([]int, len(dice))
		copy(copySlice, dice)
		sort.Ints(copySlice)

		if copySlice[0] == 1 && copySlice[1] == 2 && copySlice[2] == 3 && copySlice[3] == 4 && copySlice[4] == 5 {
			return 30
		}
		return 0
	case "big straight" :
		copySlice := make([]int, len(dice))
		copy(copySlice, dice)
		sort.Ints(copySlice)
		if copySlice[0] == 2 && copySlice[1] == 3 && copySlice[2] == 4 && copySlice[3] == 5 && copySlice[4] == 6 {
			return 30
		}
		return 0
	case "choice":
		for _, die := range dice {
			score += die
		}
		return score 
	case "yacht":
		if countElements(dice[0], dice) == 5 {
			return 50
		}
		return 0
	case "four of a kind":
		for _, die := range dice {
			if countElements(die, dice) >= 4 {
				return die * 4
			}
		}
		return 0
	case "full house":
		counts := make(map[int]int)
		for _, die := range dice {
			counts[die]++
		}
		if len(counts) == 2 {
			var n int
			for _, count := range counts {
				if count == 2 || count == 3 {
					n+=count
				}
			}
			if n == 5 {
				for die,count := range counts {
					score += die * count
				}
			}
		}
		return score
	}

		
	for _, die := range dice {
		if category == "ones" && die == 1 {
			score += 1
		}
		if category == "twos" && die == 2 {
			score += 2
		}
		if category == "threes" && die == 3 {
			score += 3
		}
		if category == "fours" && die == 4 {
			score += 4
		}
		if category == "fives" && die == 5 {
			score += 5
		}
		if category == "sixes" && die == 6 {
			score += 6
		}
	}

	return score
}


func countElements(element int,slice []int) int {
	counts := make(map[int]int)

	for _, elem := range slice {
		counts[elem]++
	}

	return counts[element]
}
