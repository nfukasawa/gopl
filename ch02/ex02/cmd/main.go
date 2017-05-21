package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/nfukasawa/gopl/ch02/ex02/conv"
)

func printTempConv(v float64, w io.Writer) {
	fmt.Fprintln(w, "Temp:")
	k := conv.Kelvin(v)
	f := conv.Fahrenheit(v)
	c := conv.Celsius(v)
	fmt.Fprintf(w, "%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
		k, conv.KToC(k), conv.KToF(k),
		f, conv.FToK(f), conv.FToC(f),
		c, conv.CToK(c), conv.CToF(c))
}

func printLengthConv(v float64, w io.Writer) {
	fmt.Fprintln(w, "Length:")
	m := conv.Meter(v)
	ft := conv.Feet(v)
	fmt.Fprintf(w, "%s = %s\n%s = %s\n", m, conv.MToF(m), ft, conv.FToM(ft))
}

func printWeightConv(v float64, w io.Writer) {
	fmt.Fprintln(w, "Weight:")
	kg := conv.Kilogram(v)
	lb := conv.Pound(v)
	fmt.Fprintf(w, "%s = %s\n%s = %s\n", kg, conv.KToP(kg), lb, conv.PToK(lb))
}

func printConv(s string) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse error:", err)
		os.Exit(1)
	}

	w := os.Stdout
	printTempConv(v, w)
	printLengthConv(v, w)
	printWeightConv(v, w)
}

func main() {
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			printConv(scanner.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			printConv(arg)
		}
	}
}
