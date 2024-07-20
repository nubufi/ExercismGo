package darts

import "math"

func Score(x, y float64) int {
	r := math.Sqrt(x*x+y*y)

	switch {
	case r>10:
		return 0
	case r>5:
		return 1
	case r>1:
		return 5
	default:
		return 10
	}
}
