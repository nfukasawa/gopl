package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type FetchResult struct {
	URI        string
	Error      error
	StatusCode int
	Body       io.ReadCloser
	StartAt    time.Time
	Duration   time.Duration
}

func Fetch(uri string) *FetchResult {
	start := time.Now()
	resp, err := http.Get(uri)
	if err != nil {
		return &FetchResult{
			URI:      uri,
			Error:    err,
			StartAt:  start,
			Duration: time.Since(start),
		}
	}

	return &FetchResult{
		URI:        uri,
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
		StartAt:    start,
		Duration:   time.Since(start),
	}
}

func FetchAll(uris []string) []*FetchResult {
	ch := make(chan *FetchResult)

	for _, uri := range uris {
		go func(uri string) {
			ch <- Fetch(uri)
		}(uri)
	}

	results := []*FetchResult{}
	for range uris {
		res := <-ch
		results = append(results, res)
	}
	return results
}

func Dump(rs []*FetchResult, w io.Writer) error {
	for _, r := range rs {
		fmt.Fprintf(w, "uri: %q\n", r.URI)
		fmt.Fprintf(w, "time: %v\n", r.StartAt)

		if r.Error != nil {
			// dump
			fmt.Fprintf(w, "error: %q\n", r.Error)

			// print
			fmt.Printf("while reading %s: %v\n", r.URI, r.Error)
			continue
		}
		defer r.Body.Close()

		// dump
		fmt.Fprintf(w, "body:\n")
		nbytes, err := io.Copy(w, r.Body)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "body_size: %d\n", nbytes)
		fmt.Fprintf(w, "duration: %f\n", r.Duration.Seconds())
		fmt.Fprintf(w, "\n")

		// print
		fmt.Printf("%.2fs  %7d  %s\n", r.Duration.Seconds(), nbytes, r.URI)
	}
	return nil
}

func main() {
	start := time.Now()
	res := FetchAll(os.Args[1:])

	file, err := os.Create("results_" + start.Format("20060102150405") + ".txt")
	if err != nil {
		fmt.Println("dump file create error: %q", err)
	}
	defer file.Close()

	if err := Dump(res, file); err != nil {
		fmt.Println("dump error: %q", err)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
