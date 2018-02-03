package v4

import (
	"math"
)

func isEven(x int) bool {
	if x&1 == 0 {
		return true
	}
	return false
}


func isPrime(n int) bool {
	if n == 0 || n == 1 || isEven(n) {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i++ {
		if isEven(i) {
			continue
		}
		if math.Mod(float64(n), float64(i)) == 0.0 {
			return false
		}
	}
	return true
}

func LargestPrimeFactor(n int) int {
	biggest := 0
	for i := 0; i < n; i++ {
		if isEven(i) {
			continue
		}
		if math.Mod(float64(n), float64(i)) == 0.0 { // i divides n evenly
			if isPrime(i) {
				if i > biggest {
					biggest = i
				}
			}
		}
	}
	return biggest
}