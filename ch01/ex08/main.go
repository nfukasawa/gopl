package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type fetchErrorKind int

const (
	fetchErrorKindNone fetchErrorKind = iota
	fetchErrorKindHTTP
	fetchErrorKindCopy
)

type fetchError struct {
	cause error
	kind  fetchErrorKind
}

func (e fetchError) Error() string {
	return e.cause.Error()
}

func IsFetchHTTPError(err error) bool {
	if e := toFetchError(err); e != nil {
		return e.kind == fetchErrorKindHTTP
	}
	return false
}

func IsFetchCopyError(err error) bool {
	if e := toFetchError(err); e != nil {
		return e.kind == fetchErrorKindCopy
	}
	return false
}

func toFetchError(err error) *fetchError {
	if e, ok := err.(*fetchError); ok {
		return e
	}
	return nil
}

func Fetch(uri string, w io.Writer) error {
	uri = complementURI(uri)
	resp, err := http.Get(uri)
	if err != nil {
		return &fetchError{cause: err, kind: fetchErrorKindHTTP}
	}
	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return &fetchError{cause: err, kind: fetchErrorKindCopy}
	}
	return nil
}

func complementURI(u string) string {
	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		return "http://" + u
	}
	return u
}

func main() {
	for _, u := range os.Args[1:] {
		err := Fetch(u, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
	}
}
