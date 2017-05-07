package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Echo struct {
	writer io.Writer
}

func NewEcho(w io.Writer) *Echo {
	return &Echo{writer: w}
}

func (e Echo) Run(args []string) {
	fmt.Fprintln(e.writer, strings.Join(args, " "))
}

func main() {
	echo := NewEcho(os.Stdout)
	echo.Run(os.Args)
}
