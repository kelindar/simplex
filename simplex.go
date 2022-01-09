package simplex

import "math"

var (
	perm   []uint8
	perm12 []uint8
	f2     = float32(0.5 * (math.Sqrt(3) - 1))
	g2     = float32((3 - math.Sqrt(3)) / 6)
)

var table = []uint8{151, 160, 137, 91, 90, 15,
	131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
	190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
	88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
	77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
	102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
	135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
	5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
	223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
	129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
	251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
	49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
	138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180}

var gradients2D = []int8{
	1, 1, -1, 1, 1, -1, -1, -1,
	1, 0, -1, 0, 1, 0, -1, 0,
	0, 1, 0, -1, 0, 1, 0, -1,
}

func init() {
	for i := 0; i < 512; i++ {
		perm = append(perm, table[i&255])
		perm12 = append(perm12, perm[i]%12)
	}
}

// Noise2 computes a two dimensional simplex noise
// Public Domain: https://weber.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf
// Reference: https://mrl.cs.nyu.edu/~perlin/noise/
func Noise2(x, y float32) float32 {

	// Skew the input space to determine which simplex cell we're in
	s := (x + y) * f2
	i := floor(x + s)
	j := floor(y + s)

	t := float32(i+j) * g2
	X0 := float32(i) - t // Unskew the cell origin back to (x,y) space
	Y0 := float32(j) - t
	x0 := x - X0 // Unskew the cell origin back to (x,y) space
	y0 := y - Y0

	// For the 2D case, the simplex shape is an equilateral triangle.
	// Determine which simplex we are in
	i1, j1 := 0, 1 // upper triangle
	if x0 > y0 {   // lower triangle
		i1 = 1
		j1 = 0
	}

	// Offsets for middle corner in (x,y) unskewed coords
	x1 := x0 - float32(i1) + g2
	y1 := y0 - float32(j1) + g2

	// Offsets for middle corner in (x,y) unskewed coords
	x2 := x0 - 1 + 2*g2
	y2 := y0 - 1 + 2*g2

	// Work out the hashed gradient indices of the three simplex corners
	ii := i & 255
	jj := j & 255
	gi0 := perm12[ii+int(perm[jj])]
	gi1 := perm12[ii+i1+int(perm[jj+j1])]
	gi2 := perm12[ii+1+int(perm[jj+1])]

	// Calculate the contribution from the three corners
	t0 := 0.5 - x0*x0 - y0*y0
	t1 := 0.5 - x1*x1 - y1*y1
	t2 := 0.5 - x2*x2 - y2*y2

	n0 := float32(0.0)
	n1 := float32(0.0)
	n2 := float32(0.0)

	if t0 >= 0 {
		t0 *= t0
		n0 = t0 * t0 * dot2D(gi0, x0, y0)
	}

	if t1 >= 0 {
		t1 *= t1
		n1 = t1 * t1 * dot2D(gi1, x1, y1)
	}

	if t2 >= 0 {
		t2 *= t2
		n2 = t2 * t2 * dot2D(gi2, x2, y2)
	}

	// Add contributions from each corner to get the final noise value.
	// The result is scaled to return values in the interval [-1,1].
	return 70.0 * (n0 + n1 + n2)
}

// dot2D computes dot product with the gradient
func dot2D(grad uint8, x, y float32) float32 {
	return float32(gradients2D[grad])*x + float32(gradients2D[grad+1])*y
}

// floor floors the floating-point value to an integer
func floor(x float32) int {
	v := int(x)
	if x < float32(v) {
		return v - 1
	}
	return v
}
