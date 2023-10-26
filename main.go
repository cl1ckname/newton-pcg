package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

const (
	H = 800 * 2
	W = 800 * 2
)

type Unary func(complex128) complex128

func NewtonIter(f, fDx Unary, start complex128, nit int) complex128 {
	for i := 0; i < nit; i++ {
		start = start - f(start)/fDx(start)
	}
	return start
}

func Dst(p1, p2 complex128) float64 {
	dx := real(p1) - real(p2)
	dy := imag(p1) - imag(p2)
	return math.Sqrt(dx*dx + dy*dy)
}

func ClosetPoint(p complex128, ps []complex128) int {
	cls := 0
	dst := Dst(p, ps[0])
	for i, p2 := range ps {
		d := Dst(p, p2)
		if d < dst {
			dst = d
			cls = i
		}
	}
	return cls
}

func main() {
	n := 6
	GeneratePool(n, W, H)
}

func GeneratePool(n, w, h int) {
	poly := RandomPoly(n)
	polyPrime := poly.Prime()
	roots := poly.Roots()
	log.Println(n)
	for _, r := range roots {
		log.Println("root ", r)
	}
	img := make([][]float64, h, h)
	for i := 0; i < h; i++ {
		img[i] = make([]float64, w, w)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			p := NewtonIter(poly.Eval, polyPrime.Eval, complex(float64(x-W/2), float64(y-H/2)), 10)
			closestRoot := ClosetPoint(p, roots)
			img[y][x] = 255. / float64(len(roots)-closestRoot)
		}
	}

	drawAndSave(img, roots)
}

func drawAndSave(m [][]float64, roots []complex128) {
	h := len(m)
	w := len(m[0])

	im := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a := uint8(m[x][y])
			im.Set(x, y, color.RGBA{
				R: a,
				G: a,
				B: a,
				A: 255,
			})
		}
	}
	for _, r := range roots {
		x := int(real(r)) + w/2
		y := int(imag(r)) + h/2

		for i := x; i < x+10; i++ {
			for j := y; j < y+10; j++ {
				im.Set(i, j, color.RGBA{200, 0, 0, 255})
			}
		}
		//draw.Draw(im, image.Rect(x, y, x+10, y+10), image.White, image.Pt(0, 0), draw.Src)
	}

	//im.Set(400, 400, color.RGBA{255, 255, 255, 255})
	//draw.Draw(im, image.Rect(300, 300, 500, 500), image.White, image.Pt(0, 0), draw.Src)
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
