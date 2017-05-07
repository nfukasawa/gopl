package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}))
	defer mock.Close()

	buf := bytes.NewBuffer(nil)
	status, err := Fetch(strings.TrimPrefix(mock.URL, "http://"), buf)
	if err != nil {
		t.Fatalf("fetch error => %q", err)
	}
	if status != 200 {
		t.Fatalf("fetch status => %d", status)
	}
	if str := buf.String(); str != "hello" {
		t.Fatalf("fetch result => %s", str)
	}
}

func TestFetch_StatusError(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "not found")
	}))
	defer mock.Close()

	buf := bytes.NewBuffer(nil)
	status, err := Fetch(strings.TrimPrefix(mock.URL, "http://"), buf)
	if err != nil {
		t.Fatalf("fetch error => %q", err)
	}
	if status != 404 {
		t.Fatalf("fetch status => %d", status)
	}
	if str := buf.String(); str != "not found" {
		t.Fatalf("fetch result => %s", str)
	}
}

func TestFetch_HTTPError(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(nil)
	}))
	defer mock.Close()

	buf := bytes.NewBuffer(nil)
	_, err := Fetch(strings.TrimPrefix(mock.URL, "http://"), buf)
	if !IsFetchHTTPError(err) {
		t.Fatalf("error => %q, want http error", err)
	}
}

func TestFetch_CopyError(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}))
	defer mock.Close()

	buf := &errorWriterMock{}
	_, err := Fetch(strings.TrimPrefix(mock.URL, "http://"), buf)
	if !IsFetchCopyError(err) {
		t.Fatalf("error => %q, want copy error", err)
	}
}

type errorWriterMock struct{}

func (w *errorWriterMock) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write error")
}
