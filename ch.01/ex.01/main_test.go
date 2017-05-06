package main

import (
	"bytes"
	"testing"
)

type echoTest struct {
	args     []string
	expected string
}

var echoTests = []echoTest{
	{[]string{""}, "\n"},
	{[]string{"echo"}, "echo\n"},
	{[]string{"echo", "hello"}, "echo hello\n"},
}

func TestEcho(t *testing.T) {
	for _, test := range echoTests {
		out := &bytes.Buffer{}
		echo := NewEcho(out)
		echo.Run(test.args)
		if result := out.String(); result != test.expected {
			t.Fatalf("echo(%q) => %q, want: %q", test.args, result, test.expected)
		}
	}

}
