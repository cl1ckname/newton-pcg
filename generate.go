package npcg

import (
	"log"
)

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
	for _, r := range roots {
		log.Println("root ", r)
	}
	img := make([][]int, h, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w, w)
	}

	for p := range mesh(w, h) {
		x, y := p.X, p.Y
		xx := float64(x) / opts.Scale
		yy := float64(y) / opts.Scale
		p := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), opts.Nit, opts.A)
		closestRoot := ClosetPoint(p, roots)
		img[y][x] = closestRoot
	}
	return img
}
