package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

const (
	H = 800 * 2
	W = 800 * 2
)

type Unary func(complex128) complex128

func main() {
	n := 6
	poly1 := RandomPoly(n, W)
	img1 := GeneratePool(poly1, W, H, 4)

	poly2 := RandomPoly(n, W)
	img2 := GeneratePool(poly2, W, H, 16)

	for p := range mesh(W, H) {
		img1[p.Y][p.X] *= img2[p.Y][p.X] / 255
	}

	poly3 := RandomPoly(n, W)
	img3 := GeneratePool(poly3, W, H, 32)

	for p := range mesh(W, H) {
		img1[p.Y][p.X] *= img3[p.Y][p.X] / 255
	}

	drawAndSave(img1, poly1.Roots())
}

func GeneratePool(poly Poly, w, h int, scale float64) [][]float64 {
	polyPrime := poly.Prime()
	roots := poly.Roots()
	for _, r := range roots {
		log.Println("root ", r)
	}
	img := make([][]float64, h, h)
	for i := 0; i < h; i++ {
		img[i] = make([]float64, w, w)
	}

	for p := range mesh(w, h) {
		x, y := p.X, p.Y
		xx := float64(x) / scale
		yy := float64(y) / scale
		p := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), 5)
		closestRoot := ClosetPoint(p, roots)
		img[y][x] = 255. / float64(len(roots)-closestRoot)
	}
	return img
}

func drawAndSave(m [][]float64, roots []complex128) {
	h := len(m)
	w := len(m[0])

	im := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a := uint8(m[x][y])
			im.Set(x, y, HSVColor{
				H: uint16(a) * 256,
				S: 128,
				V: 255,
			})
		}
	}
	for _, r := range roots {
		x := int(real(r))
		y := int(imag(r))

		for i := x; i < x+10; i++ {
			for j := y; j < y+10; j++ {
				im.Set(i, j, color.RGBA{200, 0, 0, 255})
			}
		}
	}
	f, err := os.Create("out.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	opt := jpeg.Options{Quality: 50}
	err = jpeg.Encode(f, im, &opt)
	if err != nil {
		log.Fatal(err)
	}
}
