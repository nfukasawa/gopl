package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"

	"strconv"

	"github.com/nfukasawa/gopl/ch03/ex01-04/svg"
)

func main() {
	var funcFlag string
	var serveFlag bool
	var portFlag int
	flag.StringVar(&funcFlag, "func", "", "select function to rander.")
	flag.BoolVar(&serveFlag, "serve", false, "run http server")
	flag.IntVar(&portFlag, "port", 8989, "specify server port")
	flag.Parse()

	var fun func(x, y float64) float64
	switch funcFlag {
	case "sin":
		fun = func(x, y float64) float64 {
			r := math.Hypot(x, y)
			return math.Sin(r) / r
		}
	case "eggbox":
		fun = func(x, y float64) float64 {
			return (math.Sin(x) * math.Sin(y)) / 5
		}
	case "saddle":
		fun = func(x, y float64) float64 {
			return (math.Pow(x, 2) - math.Pow(y, 2)) / 900
		}
	default:
		flag.Usage()
		os.Exit(-1)
	}

	if serveFlag {
		serve(fun, portFlag)
		return
	}

	err := svg.SVG(os.Stdout, defaultConf(fun))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%s\n", err)
		return
	}
}

func serve(fun func(x, y float64) float64, port int) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		conf := defaultConf(fun)
		q := r.URL.Query()
		if w, err := strconv.Atoi(q.Get("width")); err == nil {
			conf.Width = w
		}
		if h, err := strconv.Atoi(q.Get("height")); err == nil {
			conf.Height = h
		}
		if maxcol, err := strconv.ParseInt(q.Get("maxcolor"), 16, 32); err == nil {
			conf.MaxColor = uint32(maxcol)
		}
		if mincol, err := strconv.ParseInt(q.Get("mincolor"), 16, 32); err == nil {
			conf.MinColor = uint32(mincol)
		}

		if err := conf.Validate(); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("invalid param"))
			return
		}

		w.Header().Set("Content-Type", "image/svg+xml")
		svg.SVG(w, conf)
	}
	http.HandleFunc("/", handler)
	fmt.Println("serving on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

func defaultConf(fun func(x, y float64) float64) *svg.Config {
	return &svg.Config{
		Width: 600, Height: 320, Cells: 100, XYRange: 30.0,
		Angle:    math.Pi / 6,
		MaxColor: 0xff0000, MinColor: 0x0000ff,
		Func: fun,
	}
}
