// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black,
	color.RGBA{255, 0, 0, 0xff}, color.RGBA{0, 255, 0, 0xff},
	color.RGBA{0, 0, 255, 0xff}, color.RGBA{192, 192, 192, 0xff},
	color.RGBA{255, 255, 0, 0xff}, color.RGBA{255, 0, 255, 0xff}}

const (
	whiteIndex  = 0 // first color in palette
	blackIndex  = 1 // next color in palette
	redIndex    = 2
	greenIndex  = 3
	blueIndex   = 4
	greyIndex   = 5
	yellowIndex = 6
	ceriseIndex = 7
)

func main() {
	// The sequence of image is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	c := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(c))
		}
		if c >= len(palette) {
			c = 0
		}
		c++
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
