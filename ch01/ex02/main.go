package main

import (
	"fmt"
	"io"
	"os"
)

type Echo2 struct {
	writer io.Writer
}

func NewEcho2(w io.Writer) *Echo2 {
	return &Echo2{writer: w}
}

func (e Echo2) Run(args []string) {
	if len(args) <= 1 {
		fmt.Fprintln(e.writer, "")
		return
	}

	for i, arg := range args[1:] {
		fmt.Fprintf(e.writer, "%d: %s\n", i+1, arg)
	}
}

func main() {
	echo := NewEcho2(os.Stdout)
	echo.Run(os.Args)
}
