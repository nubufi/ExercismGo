package allyourbase

import "errors"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	decimal := 0
	for _, digit := range inputDigits {
		decimal = decimal*inputBase + digit
	}

	if decimal == 0 {
		return []int{0}, nil
	}

	outputDigits := []int{}

	for decimal > 0 {
		outputDigits = append([]int{decimal % outputBase}, outputDigits...)
		decimal /= outputBase
	}

	return outputDigits, nil
}
