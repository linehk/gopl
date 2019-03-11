package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form err: %v", err)
	}
	x := firstOrDef(r.Form["x"], 0)
	y := firstOrDef(r.Form["y"], 0)
	zoom := firstOrDef(r.Form["zoom"], 0)
	render(w, x, y, zoom)
}

func firstOrDef(forms []string, def float64) float64 {
	if len(forms) == 0 {
		return def
	}

	v, err := strconv.ParseFloat(forms[0], 64)
	if err != nil {
		return def
	}

	return v
}

func render(out io.Writer, x, y, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	exp2 := math.Exp2(1 - zoom)
	xmin, xmax := x-exp2, x+exp2
	ymin, ymax := y-exp2, y+exp2

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
