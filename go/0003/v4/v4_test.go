package v4

import (
	"testing"
	"reflect"
)

const NUMBER int = 600851475143
const benchmarkNumber int = 13195
const benchmarkPrime int = 29

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "even, not prime", args: args{n: 8}, want: false},
		{name: "odd, not prime", args: args{n: 15}, want: false},
		{name: "zero, not prime", args: args{n: 0}, want: false},
		{name: "one, not prime", args: args{n: 1}, want: false},
		{name: "a prime number, 7", args: args{n: 7}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_LargestPrimeFactor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "benchmarkNumber", args: args{n: benchmarkNumber}, want: benchmarkPrime},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPrimeFactor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LargestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_LargestPrimeFactor(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{name: "benchmarkNumber", input: benchmarkNumber},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LargestPrimeFactor(bm.input)
			}
		})
	}
}

func Benchmark_isPrime(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{name: "benchmarkNumber", input: benchmarkNumber},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPrime(bm.input)
			}
		})
	}
}

func Benchmark_isEven(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{name: "Odd number, 3", input: 3},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isEven(bm.input)
			}
		})
	}
}

