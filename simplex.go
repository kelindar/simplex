package simplex

const (
	f2 = 0.36602542 // float32(0.5 * (math.Sqrt(3) - 1))
	g2 = 0.21132487 // float32((3 - math.Sqrt(3)) / 6)
)

var (
	perm [512]uint8
	grad [512][2]float32
)

var table = [...]uint8{151, 160, 137, 91, 90, 15,
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

func init() {
	var g2d = [12]uint16{
		0x0101, // [+1, +1]
		0xff01, // [-1, +1]
		0x01ff, // [+1, -1]
		0xffff, // [-1, -1]
		0x0100, // [+1, +0]
		0xff00, // [-1, +0]
		0x0100, // [+1, +0]
		0xff00, // [-1, +0]
		0x0001, // [+0, +1]
		0x00ff, // [+0, -1]
		0x0001, // [+0, +1]
		0x00ff, // [+0, -1]
	}

	for i := 0; i < 512; i++ {
		perm[i] = table[i&255]
		idx := g2d[perm[i]%12]
		gx := int8(idx >> 8)
		gy := int8(idx)
		grad[i] = [2]float32{float32(gx), float32(gy)}
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

	// Unskew the cell origin back to (x,y) space
	t := float32(i+j) * g2
	x0 := x - (float32(i) - t)
	y0 := y - (float32(j) - t)

	// For the 2D case, the simplex shape is an equilateral triangle.
	// Determine which simplex we are in
	i1, j1 := float32(0), float32(1) // upper triangle
	if x0 > y0 {                     // lower triangle
		i1 = 1
		j1 = 0
	}

	// Offsets for middle corner in (x,y) unskewed coords
	x1 := x0 - i1 + g2
	y1 := y0 - j1 + g2

	// Offsets for middle corner in (x,y) unskewed coords
	const g = 2*g2 - 1
	x2 := x0 + g
	y2 := y0 + g

	// Work out the hashed gradient indices of the three simplex corners
	pp := perm[j&255:]
	gg := grad[i&255:]
	p0 := int(pp[0])
	p1 := int(pp[int(j1)])
	p2 := int(pp[1])
	g0 := gg[p0]
	g1 := gg[int(i1)+p1]
	g2 := gg[1+p2]

	// Calculate the contribution from the three corners
	n := float32(0.0)
	if t := 0.5 - x0*x0 - y0*y0; t > 0 {
		n += pow4(t) * (g0[0]*x0 + g0[1]*y0)
	}
	if t := 0.5 - x1*x1 - y1*y1; t > 0 {
		n += pow4(t) * (g1[0]*x1 + g1[1]*y1)
	}
	if t := 0.5 - x2*x2 - y2*y2; t > 0 {
		n += pow4(t) * (g2[0]*x2 + g2[1]*y2)
	}

	// Add contributions from each corner to get the final noise value.
	// The result is scaled to return values in the interval [-1,1].
	return 70.0 * n
}

// pow4 lifts the value to the power of 4
func pow4(v float32) float32 {
	v *= v
	return v * v
}

// floor floors the floating-point value to an integer
func floor(x float32) int {
	v := int(x)
	if x < float32(v) {
		return v - 1
	}
	return v
}
