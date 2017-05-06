package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

type echoTest struct {
	args     []string
	expected string
}

var echoTests = []echoTest{
	{[]string{""}, "\n"},
	{[]string{"echo"}, "\n"},
	{[]string{"echo", "hello", "world"}, "hello world\n"},
}

func TestEfficientEcho(t *testing.T) {
	testEcho(t, NewEfficientEcho)
}

func TestInefficientEcho(t *testing.T) {
	testEcho(t, NewInefficientEcho)
}

func testEcho(t *testing.T, echoConstructor func(io.Writer) Echo) {
	for _, test := range echoTests {
		out := &bytes.Buffer{}
		echo := echoConstructor(out)
		echo.Run(test.args)
		if result := out.String(); result != test.expected {
			t.Fatalf("echo(%q) => %q, want: %q", test.args, result, test.expected)
		}
	}
}

func BenchmarkEfficientEcho100(b *testing.B) {
	benchEcho(b, NewEfficientEcho, 100)
}

func BenchmarkEfficientEcho10000(b *testing.B) {
	benchEcho(b, NewEfficientEcho, 10000)
}

func BenchmarkInefficientEcho100(b *testing.B) {
	benchEcho(b, NewInefficientEcho, 100)
}

func BenchmarkInefficientEcho10000(b *testing.B) {
	benchEcho(b, NewInefficientEcho, 10000)
}

func benchEcho(b *testing.B, echoConstructor func(io.Writer) Echo, n int) {
	echo := echoConstructor(ioutil.Discard)
	args := makeArgs(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		echo.Run(args)
	}
}

func makeArgs(n int) []string {
	args := make([]string, n+1)
	args[0] = "echo"
	for i := 1; i < n+1; i++ {
		args[i] = "hello"
	}
	return args
}
