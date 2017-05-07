package main

import (
	"reflect"
	"testing"
)

func TestDupMap(t *testing.T) {
	dupmap := NewDupMap()
	dupmap.CountFiles([]string{"fixture/test1.txt", "fixture/test2.txt", "fixture/test3.txt"})

	dups := dupmap.EnumDup()
	if l := len(dups); l != 4 {
		t.Fatalf("len(dupmap.EnumDup()) => %d, want: 4", l)
	}

	testDup(t, dups, "hello", 2, "fixture/test1.txt", "fixture/test2.txt")
	testDup(t, dups, "bbb", 2, "fixture/test1.txt", "fixture/test2.txt")
	testDup(t, dups, "fuga", 3, "fixture/test1.txt", "fixture/test1.txt", "fixture/test3.txt")
	testDup(t, dups, "foo", 2, "fixture/test2.txt", "fixture/test3.txt")
}

func testDup(t *testing.T, dups map[string]*Dup, txt string, count int, files ...string) {
	dup, ok := dups[txt]
	if !ok {
		t.Fatalf("dups[%s] is nil", txt)
	}

	if dup.Count != count {
		t.Fatalf("dups[%s] count => %d, want: %d", dup.Count, count)
	}

	if !reflect.DeepEqual(dup.Files, files) {
		t.Fatalf("dups[%s] files => %q, want: %q", dup.Files, files)
	}
}
