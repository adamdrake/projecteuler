package v3

import "math"

// The even check should have slowed things down, but it
// turns out we're doing more divisi


func isEven(x int) bool {
	if x&1 == 0 {
		return true
	}
	return false
}

// Ignore even factors since they can never be prime
func factors(n int) []int {
	facs := make([]int, 0)
	for i := 2; i < n; i++ {
		if isEven(i) {
			continue
		}
		if math.Mod(float64(n), float64(i)) == 0.0 {
			facs = append(facs, i)
		}
	}
	return facs
}

// Even numbers can't be prime, so we don't need to check them
func isPrime(n int) bool {
	if n == 0 || n == 1 || isEven(n) {
		return false
	}
	facs := factors(n)
	if len(facs) != 0 {
		return false
	}
	return true
}

func LargestPrimeFactor(n int) int {
	biggest := 0
	for i := 2; i < n; i++ {
		if isEven(i) { // skip even numbers
			continue
		}
		if math.Mod(float64(n), float64(i)) == 0.0 {
			if isPrime(i) {
				if i > biggest {
					biggest = i
				}
			}
		}
	}
	return biggest
}