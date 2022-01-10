package simplex

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkNoise/10x10-8         	  763042	      1568 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoise/100x100-8       	    7402	    159403 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoise/1000x1000-8     	      74	  15732020 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkNoise(b *testing.B) {
	var out float32
	for _, size := range []int{10, 100, 1000} {
		b.Run(fmt.Sprintf("2d_%vx%v", size, size), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for n := 0; n < b.N; n++ {
				for x := 0; x < size; x++ {
					for y := 0; y < size; y++ {
						out = Noise2(float32(x), float32(y))
					}
				}
			}
		})

	}

	assert.NotZero(b, out)
}

func BenchmarkDot(b *testing.B) {
	var out float32
	b.Run("2d", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			for i := uint8(0); i < 100; i++ {
				out = dot2D(i%12, 10, 20)
			}
		}
	})

	assert.NotZero(b, out)
}

func TestSimplex_500x500(t *testing.T) {
	n := 500
	freq := float32(25)
	img := image.NewGray(image.Rect(0, 0, n, n))
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			v := (1 + Noise2(float32(x)/freq, float32(y)/freq)) / 2
			img.Set(x, y, color.Gray{
				Y: uint8(v * 255),
			})
		}
	}

	// Compare with the reference
	expect, err := os.Open("fixtures/500.png")
	assert.NoError(t, err)
	out, err := png.Decode(expect)
	assert.NoError(t, err)
	assert.Equal(t, out, img)

	//f, err := os.Create("out.png")
	//assert.NoError(t, err)
	//assert.NoError(t, png.Encode(f, img))
}

func TestFloor(t *testing.T) {
	assert.Equal(t, int(math.Floor(1.5)), floor(1.5))
	assert.Equal(t, int(math.Floor(0.5)), floor(0.5))
	assert.Equal(t, int(math.Floor(-1.5)), floor(-1.5))
}
