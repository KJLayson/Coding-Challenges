package prime

import (
	"math/big"
)

func isPrime(n int64) bool {
	big_int := big.NewInt(n)
	f_prime := big_int.ProbablyPrime(20)

	return f_prime
}

func isCompleteDivision(number int64, prime int64) bool {
	return number%prime == 0
}

func Factors(n int64) []int64 {
	var list_factors = []int64{}
	if n <= 1 {
		return list_factors
	}

	for i := int64(2); i <= n; i++ {
		if isPrime(i) {
			if isCompleteDivision(n, i) {
				list_factors = append(list_factors, i)
				n /= i
				i = 1
			}
		}
	}

	return list_factors
}
