package core

type Unary func(complex128) complex128

type GenerationOpts struct {
	Scale  float64
	A      complex128
	Offset complex128
	Nit    int
}

func RandomPool(n, W, H int, opts GenerationOpts) [][]int {
	poly1 := RandomPoly(n, opts.Scale)
	return GeneratePool(poly1, W, H, opts)
}

func GeneratePool(poly Poly, w, h int, opts GenerationOpts) [][]int {
	polyPrime := poly.Prime()
	roots := poly.Roots()
	img := make([][]int, h, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w, w)
	}

	for p := range Mesh(w, h) {
		x, y := p.X, p.Y
		xx := (float64(x) + real(opts.Offset)) / opts.Scale
		yy := (float64(y) + imag(opts.Offset)) / opts.Scale
		p := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), opts.Nit, opts.A)
		closestRoot := ClosetPoint(p, roots)
		img[y][x] = closestRoot
	}
	return img
}
