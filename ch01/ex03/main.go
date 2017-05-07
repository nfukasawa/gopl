package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Echo interface {
	Run(args []string)
}

type EfficientEcho struct {
	writer io.Writer
}

func NewEfficientEcho(w io.Writer) Echo {
	return &EfficientEcho{writer: w}
}

func (e EfficientEcho) Run(args []string) {
	fmt.Fprintln(e.writer, strings.Join(args[1:], " "))
}

type InefficientEcho struct {
	writer io.Writer
}

func NewInefficientEcho(w io.Writer) Echo {
	return &InefficientEcho{writer: w}
}

func (e InefficientEcho) Run(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(e.writer, s)
}

func main() {
	echo := NewEfficientEcho(os.Stdout)
	echo.Run(os.Args)
}
