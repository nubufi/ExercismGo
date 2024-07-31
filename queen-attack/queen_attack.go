package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, errors.New("same square")
	}
	if !checkPosition(whitePosition) || !checkPosition(blackPosition) {
		return false, errors.New("invalid position")
	}
	if whitePosition[0] == blackPosition[0] || whitePosition[1] == blackPosition[1] {
		return true, nil
	}
	if math.Abs(float64(whitePosition[0])-float64(blackPosition[0])) == math.Abs(float64(whitePosition[1])-float64(blackPosition[1])) {
		return true, nil
	}
	return false, nil
}

func checkPosition(position string) bool {
	if len(position) != 2 {
		return false
	}
	if position[0] < 'a' || position[0] > 'h' {
		return false
	}
	if position[1] < '1' || position[1] > '8' {
		return false
	}
	return true
}
