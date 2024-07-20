package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	asString := strconv.Itoa(n)

	var sum float64
	for i, s := range asString {
		digit, _ := strconv.Atoi(string(s))
		sum += math.Pow(float64(digit), float64(len(asString)-i))
	}

	return int(sum) == n
}
