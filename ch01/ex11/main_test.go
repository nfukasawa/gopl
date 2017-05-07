package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	mock1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}))
	defer mock1.Close()

	mock2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "world")
	}))
	defer mock2.Close()

	mock3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(nil)
	}))
	defer mock3.Close()

	res := FetchAll([]string{mock1.URL, mock2.URL, mock3.URL})
	if len(res) != 3 {
		t.Fatalf("FetchResults length => %d, want: %d", len(res), 3)
	}

	for _, r := range res {
		switch r.URI {
		case mock1.URL:
			body, _ := ioutil.ReadAll(r.Body)
			if string(body) != "hello" {
				t.Fatalf("fetch(%q) => %q, want: %q", r.URI, string(body), "hello")
			}
		case mock2.URL:
			body, _ := ioutil.ReadAll(r.Body)
			if string(body) != "world" {
				t.Fatalf("fetch(%q) => %q, want: %q", r.URI, string(body), "world")
			}
		case mock3.URL:
			if r.Error == nil {
				t.Fatalf("fetch(%q) result should be error", r.URI)
			}
		default:
			t.Fatalf("invalid url: %q", r.URI)
		}
	}
}
