package resistorcolortrio

import (
	"fmt"
	"math"
)

var Colors = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	val := (Colors[colors[0]]*10 + Colors[colors[1]]) * int(math.Pow(10, float64(Colors[colors[2]])))
	if val == 0 {
		return "0 ohms"
	}

	unit := "ohms"
	if val >= 1e9 {
		val /= 1e9
		unit = "gigaohms"
	}

	if val >= 1e6 {
		val /= 1e6
		unit = "megaohms"
	}

	if val >= 1e3 {
		val /= 1e3
		unit = "kiloohms"
	}

	return fmt.Sprintf("%d %s", int(val), unit)

}
