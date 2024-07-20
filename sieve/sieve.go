package sieve


func Sieve(limit int) []int {

	if limit < 2 {
		return []int{}
	}

	primes := make([]int, 0, limit)

	// Create a slice of bools with length limit+1
	// Each index represents a number from 0 to limit
	// The value at each index is true if the number is prime
	// and false if the number is not prime
	isPrime := make([]bool, limit+1)

	// Initialize all numbers from 2 to limit as prime
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	// Iterate over all numbers from 2 to limit
	// If a number is prime, add it to the primes slice
	// and mark all multiples of the number as not prime
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	return primes
}
