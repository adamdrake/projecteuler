package v5

import (
	"math"
)


func isPrime(n int) bool {
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		if math.Mod(float64(n), float64(i)) == 0.0 {
			return false
		}
	}
	return true
}

func LargestPrimeFactor(n int) int {
	biggest := 0
	// First, find factors
	for i := 3; i < n; i = i + 2 {
		if math.Mod(float64(n), float64(i)) == 0.0 {
			// i is a factor, so check if it's prime
			if isPrime(i) {
				if i > biggest {
					biggest = i
				}
			}
		}
	}
	return biggest
}