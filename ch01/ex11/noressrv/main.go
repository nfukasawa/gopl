package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

const port = 8989

func main() {
	ctx := context.Background()

	var wg sync.WaitGroup
	done := make(chan struct{})

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/stop") {
				close(done)
				return
			}
			wg.Add(1)
			<-done
			wg.Done()
		}),
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	defer srv.Shutdown(ctx)

	<-done
	wg.Wait()
}
