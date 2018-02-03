package main

import "testing"

func TestArithSeq(t *testing.T) {
	if arithSeq(9) != 23 {
		t.Fatal("The sum of values from 1 to 10 which are divisible by 5 and 3 should be 23")
	}
	if arithSeq(999) != 233168 {
		t.Fatal("The sum of values from 1 to 1000 which are divisible by 5 and 3 should be 233168")
	}
}

func TestBruteForce(t *testing.T) {
	if bruteForce(9) != 23 {
		t.Fatal("The sum of values from 1 to 10 which are divisible by 5 and 3 should be 23")
	}
	if bruteForce(999) != 233168 {
		t.Fatal("The sum of values from 1 to 1000 which are divisible by 5 and 3 should be 233168")
	}
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bruteForce(999)
	}
}

func BenchmarkArithSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arithSeq(999)
	}
}
