package popcount_test

import (
	"testing"

	"github.com/nfukasawa/gopl/ch02/ex03-05/popcount"
)

const val = 0x1234567890ABCDEF

func TestPopCount(t *testing.T) {
	c := popcount.PopCount(val)

	if c0 := popcount.PopCountByLoop(val); c0 != c {
		t.Fatalf("popcount.PopCountByLoop(0x%x) => %v, want: %v", val, c0, c)
	}

	if c0 := popcount.PopCountByBitShift(val); c0 != c {
		t.Fatalf("popcount.PopCountByBitShift(0x%x) => %v, want: %v", val, c0, c)
	}

	if c0 := popcount.PopCountByBitClear(val); c0 != c {
		t.Fatalf("popcount.PopCountByBitClear(0x%x) => %v, want: %v", val, c0, c)
	}

}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(val)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByLoop(val)
	}
}

func BenchmarkPopCountByBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByBitShift(val)
	}
}

func BenchmarkPopCountByBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByBitClear(val)
	}
}
