package npcg

import (
	"log"
)

type Unary func(complex128) complex128

func RandomPool(n, W, H int, scale float64) [][]int {
	poly1 := RandomPoly(n, scale)
	return GeneratePool(poly1, W, H, scale)
}

func GeneratePool(poly Poly, w, h int, scale float64) [][]int {
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
		xx := float64(x) / scale
		yy := float64(y) / scale
		p := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), 20)
		closestRoot := ClosetPoint(p, roots)
		img[y][x] = closestRoot
	}
	return img
}
