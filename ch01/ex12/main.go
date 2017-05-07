package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	backgroundIndex = 0
	greenIndex      = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	port := 8989
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil {
			port = n
		}
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, newLissajousOptions(r.URL.Query()))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
	return
}

type lissajousOptions struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func newLissajousOptions(params url.Values) *lissajousOptions {
	return &lissajousOptions{
		cycles:  getParamInt(params, "cycles", 5),
		res:     getParamFloat(params, "res", 0.001),
		size:    getParamInt(params, "size", 100),
		nframes: getParamInt(params, "nframes", 64),
		delay:   getParamInt(params, "delay", 8),
	}

}

func getParamInt(params url.Values, name string, def int) int {
	val := params.Get(name)
	if val == "" {
		return def
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		fmt.Fprintf(os.Stderr, "param parse error %q: %q", name, val)
		return def
	}
	return n
}

func getParamFloat(params url.Values, name string, def float64) float64 {
	val := params.Get(name)
	if val == "" {
		return def
	}
	n, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "param parse error %q: %q", name, val)
		return def
	}
	return n
}

func lissajous(out io.Writer, opts *lissajousOptions) {

	cycles := opts.cycles
	res := opts.res
	size := opts.size
	nframes := opts.nframes
	delay := opts.delay

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
