package main

import (
	"log"
)

const (
	H = 160
	W = 240
)

type Unary func(complex128) complex128

func main() {
	n := 6
	poly1 := RandomPoly(n, W)
	img1 := GeneratePool(poly1, W, H, 1)

	poly2 := RandomPoly(n, W)
	img2 := GeneratePool(poly2, W, H, 2)

	for p := range mesh(W, H) {
		img1[p.Y][p.X] = (img1[p.Y][p.X] + img2[p.Y][p.X]) % n
	}

	poly3 := RandomPoly(n, W)
	img3 := GeneratePool(poly3, W, H, 4)

	for p := range mesh(W, H) {
		img1[p.Y][p.X] = (img1[p.Y][p.X] + img3[p.Y][p.X]) % n
	}

	//drawAndSave(img1, poly1.Roots())
	generateCave(img1)
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
		p := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), 5)
		closestRoot := ClosetPoint(p, roots)
		img[y][x] = closestRoot
	}
	return img
}
