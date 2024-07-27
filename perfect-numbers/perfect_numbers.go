package perfect

import (
	"errors"
)

// Define the Classification type here.
type Classification int

// Define the Classification constants here.
const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive error = errors.New("only positive")

func Classify(n int64) (Classification, error) {
	var sum int64
	
	if n <= 0 {
		return ClassificationAbundant,ErrOnlyPositive
	}
	for i:=int64(1);i<=n/2;i++ {
		if n%i == 0 {
			sum+=i
		}
	}
	
	switch  {
		case sum == n:
			return ClassificationPerfect,nil
		case sum > n:
			return ClassificationAbundant,nil
		default:
			return ClassificationDeficient,nil
	}

}
