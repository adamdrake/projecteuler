package main

import (
	"fmt"
	"math"
)

func bruteForce(limit int) int {
	total := 0
	for i := 0; i <= limit; i++ {
		if math.Mod(float64(i), 5) == 0 || math.Mod(float64(i), 3) == 0 {
			total += i
		}
	}
	return total
}

// The sum of a finite arithmetic sequence works for all real numbers
func arithSum(n float64) float64 {
	return (n * (n + 1)) / 2
}

func arithSeq(limit float64) int {
	threes := 3 * arithSum(math.Floor(limit/3))
	fives := 5 * arithSum(math.Floor(limit/5))
	fifteens := 15 * arithSum(math.Floor(limit/15))
	return int(threes + fives - fifteens)
}

func main() {
	fmt.Println(bruteForce(999))
	fmt.Println(arithSeq(9))
	fmt.Println(arithSeq(999))
}
