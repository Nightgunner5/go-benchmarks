package test

import (
	"testing"
)

func chain(in <-chan int, out chan<- int) {
	defer close(out)

	for i := range in {
		out <- i + 1
	}
}

func makeChain(length int) (in chan<- int, out <-chan int) {
	if length <= 0 {
		panic("invalid length")
	}

	ch := make(chan int)
	in = ch
	for i := 0; i < length; i++ {
		ch1 := make(chan int)
		go chain(ch, ch1)
		ch, out = ch1, ch1
	}
	return
}

func BenchmarkChain1(b *testing.B) {
	chainBench(b, 1)
}

func BenchmarkChain10(b *testing.B) {
	chainBench(b, 10)
}

func BenchmarkChain100(b *testing.B) {
	chainBench(b, 100)
}

func BenchmarkChain1000(b *testing.B) {
	chainBench(b, 1000)
}

func BenchmarkChain10000(b *testing.B) {
	chainBench(b, 10000)
}

func BenchmarkChain100000(b *testing.B) {
	chainBench(b, 100000)
}

func chainBench(b *testing.B, length int) {
	b.StopTimer()
	in, out := makeChain(length)
	b.StartTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			in <- 0
		}
		close(in)
	}()

	for i := 0; i < b.N; i++ {
		j := <-out
		if j != length {
			b.Errorf("Unexpected value at end of chain: expected %d, got %d", length, j)
		}
	}
}
