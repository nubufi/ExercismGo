package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("invalid number")
	}

	if n==1 {
		return 2,nil
	}

	primes := []int{2}
	number := 3
	for len(primes) != n {
		if isPrime(number) {
			primes = append(primes,number)
		}
		number += 2
	}

	return primes[n-1],nil
}

func isPrime(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	for i:=3;i <= sqrt;i+=2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
