package prime

import (
	"errors"
	"math/big"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("number must be greater than 0")
	}

	primes := []int{}

	for i := 1; len(primes) < n; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			primes = append(primes, i)
		}
	}
	return primes[len(primes)-1], nil
}
