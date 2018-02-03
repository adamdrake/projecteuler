package v2

import (
	"testing"
	"reflect"
)

const NUMBER int = 600851475143
const benchmarkNumber int = 13195
const benchmarkPrime int = 29

func Test_factors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "factors of zero", args: args{n: 0}, want: []int{}},
		{name: "factors of 6", args: args{n: 6}, want: []int{3}},
		{name: "factors of a prime, 7", args: args{n: 7}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factors() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Benchmark_factors(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{name: "benchmarkNumber", input: benchmarkNumber},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				factors(bm.input)
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

func Test_isEven(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Even number, 4", args: args{n: 4}, want: true},
		{name: "Even number, 0", args: args{n: 0}, want: true},
		{name: "Even number, -2", args: args{n: -2}, want: true},
		{name: "Odd number, 3", args: args{n: 3}, want: false},
		{name: "Odd number, 1", args: args{n: 1}, want: false},
		{name: "Odd number, -5", args: args{n: -5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
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