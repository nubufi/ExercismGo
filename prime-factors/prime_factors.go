package prime

func Factors(n int64) []int64 {
	var factors []int64
	if n < 2 {
		return factors
	}
	for n > 1 {
		for n%2 == 0 {
			factors = append(factors, 2)
			n /= 2
		}
		for i := 3; i <= int(n); i++ {
			for n%int64(i) == 0 {
				factors = append(factors, int64(i))
				n /= int64(i)
			}
		}
	}

	return factors

}
