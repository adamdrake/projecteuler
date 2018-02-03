package main

import (
	"fmt"
	"math"
)

func isEven(x int) bool {
	if x&1 == 0 {
		return true
	}
	return false
}

func sumEvens(xs []int) int {
	sum := 0
	for _, v := range xs {
		if isEven(v) {
			sum += v
		}
	}
	return sum
}

func sumEvenFibsSlice(limit int) int {
	fibs := []int{1, 2}
	newElement := 0
	for i := 2; newElement < limit; i++ {
		newElement = fibs[i-1] + fibs[i-2]
		fibs = append(fibs, newElement)
	}
	return sumEvens(fibs)
}

func fibsRecursion(limit int) int {
	if limit <= 2 {
		return limit
	}
	return fibsRecursion(limit-2) + fibsRecursion(limit-1)
}

func sumEvenFibsRecursion(limit int) int {
	sum := 0
	next := 0
	for i := 1; next < limit; i++ {
		next = fibsRecursion(i)
		if isEven(next) {
			sum += next
		}
	}
	return sum
}

func sumEvenFibsLoop(limit int) int {

	sum := 0
	for i, j := 1, 1; j < limit; i, j = i+j, i {
		if i&1 == 0 {
			sum += i
		}
	}
	return sum
}

func fibsChannel(limit int) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for i, j := 1, 1; j < limit; i, j = i+j, i {
			output <- i
		}
	}()
	return output
}

func sumEvenFibsChannel(limit int) int {
	sum := 0
	fibs := fibsChannel(limit)
	for f := range fibs {
		if isEven(f) {
			sum += f
		}
	}
	return sum
}

func fibApprox(n int) int {
	fib := (1 / math.Sqrt(5)) * math.Pow(math.Phi, float64(n+1))
	return int(fib + 0.5)
}

func sumEvenFibsApprox(limit int) int {
	fib := 0
	sum := 0
	for i := 2; fib < limit; i++ {
		fib = fibApprox(i)
		if isEven(fib) {
			sum += fib
		}
	}
	return sum
}

func main() {
	LIMIT := 3999999
	fmt.Println(sumEvenFibsLoop(LIMIT))
	fmt.Println(sumEvenFibsSlice(LIMIT))
	fmt.Println(sumEvenFibsRecursion(LIMIT))
	fmt.Println(sumEvenFibsChannel(LIMIT))
	fmt.Println("approx", sumEvenFibsApprox(LIMIT))
}
