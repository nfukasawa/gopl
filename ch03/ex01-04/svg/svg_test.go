package svg_test

import (
	"bytes"
	"math"
	"strings"
	"testing"

	"github.com/nfukasawa/gopl/ch03/ex01-04/svg"
)

func TestSVG_IgnoreInfoOrNaN(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	err := svg.SVG(buf, &svg.Config{
		Width: 10, Height: 10, Cells: 3, XYRange: 30.0, Angle: math.Pi / 6,
		Func: func(x, y float64) float64 {
			if x > 0 {
				return math.Inf(1)
			}
			return math.NaN()
		},
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if strings.Contains(buf.String(), "polygon") {
		t.Fatalf("Inf or NaN should be ignored")
	}
}

// TODO test for Calc
