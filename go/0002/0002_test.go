package main

import "testing"

const LIMIT int = 3999999
const SUM_LIMIT int = 4613732

func TestIsEven(t *testing.T) {
	if (isEven(4) != true) || (isEven(3) != false) {
		t.Fatal("isEven returned incorrect value")
	}
}

func BenchmarkIsEvenEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEven(2)
	}
}

func BenchmarkIsEvenOdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEven(3)
	}
}

func TestSumEvens(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := sumEvens(nums)
	if sum != 30 {
		t.Fatal("The sum of the even numbers in 0..10 inclusive is 30 but got", sum)
	}
}

func TestSumEvenFibsSlice(t *testing.T) {
	if sumEvenFibsSlice(LIMIT) != SUM_LIMIT {
		t.Fatal("")
	}
}

func BenchmarkSumEvenFibsSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumEvenFibsSlice(LIMIT)
	}
}

func TestFibsRecursion(t *testing.T) {
	firstTen := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	for i, v := range firstTen {
		fib := fibsRecursion(i + 1)
		if fib != v {
			t.Fatal("recursion gave", fib, "want", v)
		}
	}
}

func TestSumEvenFibsRecursion(t *testing.T) {
	result := sumEvenFibsRecursion(LIMIT)
	if result != SUM_LIMIT {
		t.Fatal("wanted", SUM_LIMIT, "but received", result)
	}
}

func BenchmarkSumEvenFibsRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumEvenFibsRecursion(LIMIT)
	}
}

func TestSumEvenFibsLoop(t *testing.T) {
	if sumEvenFibsLoop(LIMIT) != SUM_LIMIT {
		t.Fatal("sumevenfibloop")
	}

}

func BenchmarkSumEvenFibsLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumEvenFibsLoop(LIMIT)
	}
}

func TestFibsChannel(t *testing.T) {
	firstTen := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	fibs := fibsChannel(90)
	for _, v := range firstTen {
		chanVal := <-fibs
		if v != chanVal {
			t.Fatal("fibChannel incorrect.  Expected", v, "but received", chanVal)
		}
	}
}

func TestSumEvenFibsChannel(t *testing.T) {
	if sumEvenFibsChannel(LIMIT) != SUM_LIMIT {
		t.Fatal("")
	}
}

func BenchmarkSumEvenFibsChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumEvenFibsChannel(LIMIT)
	}
}

func TestFibApprox(t *testing.T) {
	firstTen := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	for i, v := range firstTen {
		fib := fibApprox(i + 1)
		if fib != v {
			t.Fatal("got", fib, "want", v)
		}
	}
}

func TestSumEvenFibsApprox(t *testing.T) {
	if sumEvenFibsApprox(LIMIT) != SUM_LIMIT {
		t.Fatal("")
	}
}

func BenchmarkSumEvenFibsApprox(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sumEvenFibsApprox(LIMIT)
	}
}
