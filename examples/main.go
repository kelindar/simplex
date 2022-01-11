package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/kelindar/simplex"
)

type Octave struct {
	Frequency, Scale float32
}

var octaves = []Octave{
	{.0025, .6},
	{.0001, .1},
	{.015, .2},
	{.1, .06},
	{.3, .03},
	{.9, .03},
}

var palette = []color.RGBA{
	{41, 128, 185, 255},
	{52, 152, 219, 255},
	{255, 234, 167, 255},
	{253, 203, 110, 255},
	{248, 194, 145, 255},
	{184, 233, 148, 255},
	{120, 224, 143, 255},
	{189, 195, 199, 255},
	{236, 240, 241, 255},
}

func main() {
	n := 800
	img := image.NewRGBA(image.Rect(0, 0, n, n))
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v := float32(0)
			for _, o := range octaves {
				n := (1 + simplex.Noise2(o.Frequency*(float32(x)), o.Frequency*(float32(y)))) / 2
				v += o.Scale * n
			}

			// Circular distance from the center point
			dx := float64(x)/float64(n) - 0.5
			dy := float64(y)/float64(n) - 0.5
			d := math.Sqrt(dx*dx+dy*dy) * 2
			d = math.Pow(d, 1.5)
			v = (1 - float32(d) + v) / 2

			// Squish the corners closer
			v = float32(math.Pow(float64(v), .6))
			img.Set(x, y, colorFor(v))
		}
	}

	f, _ := os.Create("terrain.png")
	png.Encode(f, img)
}

func colorFor(v float32) color.Color {
	const cutoff = .5
	switch {
	case v > cutoff:
		count := len(palette) - 1
		norm := float64(v-cutoff) / (1 - cutoff)
		bracket := int(math.Floor(norm * float64(count)))
		return palette[bracket+1]
	default:
		return palette[0]
	}
}
