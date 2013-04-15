package test

import (
	"testing"
)

func generate(ch chan<- int, max int) {
	defer close(ch)
	for i := 2; i <= max; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	defer close(out)
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}

const (
	Prime1 = 2 // 2 is the 1st prime.
	Prime10 = 29 // 29 is the 10th prime.
	Prime100 = 541 // 541 is the 100th prime.
	Prime1000 = 7919 // 7919 is the 1000th prime.
	Prime10000 = 104729 // 104729 is the 10000th prime.
	Prime100000 = 1299709 // 1299709 is the 100000th prime.
)

func BenchmarkPrimes1(b *testing.B) {
	primes(b, 1, Prime1)
}

func BenchmarkPrimes10(b *testing.B) {
	primes(b, 10, Prime10)
}

func BenchmarkPrimes100(b *testing.B) {
	primes(b, 100, Prime100)
}

func BenchmarkPrimes1000(b *testing.B) {
	primes(b, 1000, Prime1000)
}

func BenchmarkPrimes10000(b *testing.B) {
	primes(b, 10000, Prime10000)
}

func BenchmarkPrimes100000(b *testing.B) {
	primes(b, 100000, Prime100000)
}

func primes(b *testing.B, count, last int) {
	for i := 0; i < b.N; i++ {
		var prime int
		in := make(chan int)
		go generate(in, last)
		for i := 0; i < count; i++ {
			prime = <-in
			out := make(chan int)
			go filter(in, out, prime)
			in = out
		}
		if prime != last {
			b.Fatalf("Expected last prime to be %d, but it is %d.", last, prime)
		}
	}
}
