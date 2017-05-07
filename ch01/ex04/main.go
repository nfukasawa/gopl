package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dup struct {
	Count int
	Files []string
}

type DupMap map[string]*Dup

func NewDupMap() DupMap {
	return DupMap{}
}

func (d DupMap) CountFiles(files []string) error {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := d.Count(f); err != nil {
			return err
		}
	}
	return nil
}

func (d DupMap) Count(f *os.File) error {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		if dup, ok := d[txt]; ok {
			dup.Files = append(dup.Files, f.Name())
			dup.Count++
		} else {
			d[txt] = &Dup{
				Files: []string{f.Name()},
				Count: 1,
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	return nil
}
func (d DupMap) EnumDup() map[string]*Dup {
	dups := map[string]*Dup{}
	for txt, dup := range d {
		if dup.Count > 1 {
			dups[txt] = dup
		}
	}
	return dups
}

func main() {
	dupmap := NewDupMap()
	files := os.Args[1:]

	var err error
	if len(files) == 0 {
		err = dupmap.Count(os.Stdin)
	} else {
		err = dupmap.CountFiles(files)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", err)
	}

	for txt, dup := range dupmap.EnumDup() {
		fmt.Printf("%d\t%s\t%q\n", dup.Count, txt, dup.Files)
	}
}
