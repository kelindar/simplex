package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"

	"github.com/kelindar/bench"
	"github.com/kelindar/simplex"
)

var sizes = []int{1e3, 1e6}

func main() {
	bench.Run(func(b *bench.B) {
		runNoise(b)
	}, bench.WithDuration(10*time.Millisecond), bench.WithSamples(100))
}

func runNoise(b *bench.B) {
	shapes := []struct {
		name string
		gen  func(int) [][2]float32
	}{
		{"seq", dataSeq},
		{"rnd", dataRand},
		{"circ", dataCircle},
	}

	const size = 1000
	for _, shape := range shapes {
		points := shape.gen(size)
		name := fmt.Sprintf("noise %s (%s)", formatSize(size), shape.name)
		b.Run(name, func(i int) {
			p := points[i%len(points)]
			_ = simplex.Noise2(p[0], p[1])
		})
	}
}

func formatSize(size int) string {
	if size >= 1e6 {
		return fmt.Sprintf("%.0fM", float64(size)/1e6)
	}
	return fmt.Sprintf("%.0fK", float64(size)/1e3)
}

func dataSeq(n int) [][2]float32 {
	pts := make([][2]float32, n)
	for i := 0; i < n; i++ {
		f := float32(i)
		pts[i] = [2]float32{f, f}
	}
	return pts
}

func dataRand(n int) [][2]float32 {
	pts := make([][2]float32, n)
	for i := 0; i < n; i++ {
		pts[i] = [2]float32{rand.Float32(), rand.Float32()}
	}
	return pts
}

func dataCircle(n int) [][2]float32 {
	pts := make([][2]float32, n)
	for i := 0; i < n; i++ {
		angle := 2 * math.Pi * float64(i) / float64(n)
		pts[i] = [2]float32{float32(math.Cos(angle)), float32(math.Sin(angle))}
	}
	return pts
}
