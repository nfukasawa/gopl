package tempconv_test

import (
	"math"
	"testing"

	"github.com/nfukasawa/gopl/ch02/ex01/tempconv"
)

var temps = []struct {
	c tempconv.Celsius
	f tempconv.Fahrenheit
	k tempconv.Kelvin
}{
	{-273.15, -459.67, 0},
	{0, 32, 273.15},
	{100, 212, 373.15},
}

type f float64

func eq(a, b f) bool {
	const epsilon = 1e-10
	a0, b0 := float64(a), float64(b)
	return (math.Abs(a0-b0) <= epsilon) || (math.Abs(a0-b0) <= math.Max(math.Abs(a0), math.Abs(b0))*epsilon)
}

func TestCToF(t *testing.T) {
	for _, temp := range temps {
		if t0 := tempconv.CToF(temp.c); !eq(f(t0), f(temp.f)) {
			t.Fatalf("CToF(%v) => %v, want %v", temp.c, t0, temp.f)
		}
	}
}

func TestCToK(t *testing.T) {
	for _, temp := range temps {
		if t0 := tempconv.CToK(temp.c); !eq(f(t0), f(temp.k)) {
			t.Fatalf("CToK(%v) => %v, want %v", temp.c, t0, temp.k)
		}
	}
}

func TestFToC(t *testing.T) {
	for _, temp := range temps {
		if t0 := tempconv.FToC(temp.f); !eq(f(t0), f(temp.c)) {
			t.Fatalf("FToC(%v) => %v, want %v", temp.f, t0, temp.c)
		}
	}
}

func TestFToK(t *testing.T) {
	for _, temp := range temps {
		if t0 := tempconv.FToK(temp.f); !eq(f(t0), f(temp.k)) {
			t.Fatalf("FToK(%v) => %v, want %v", temp.f, t0, temp.k)
		}
	}
}

func TestKToC(t *testing.T) {
	for _, temp := range temps {
		if t0 := tempconv.KToC(temp.k); !eq(f(t0), f(temp.c)) {
			t.Fatalf("KToC(%v) => %v, want %v", temp.k, t0, temp.c)
		}
	}
}

func TestKToF(t *testing.T) {
	for _, temp := range temps {
		if t0 := tempconv.KToF(temp.k); !eq(f(t0), f(temp.f)) {
			t.Fatalf("KToF(%v) => %v, want %v", temp.k, t0, temp.f)
		}
	}
}
