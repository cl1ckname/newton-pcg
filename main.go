package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"math/cmplx"
	"os"
)

const (
	H = 800
	W = 800
)

type Unary func(complex128) complex128

func Hyper(z complex128) complex128 {
	return z*z*z - complex(100, 0)
}

func HyperDx(z complex128) complex128 {
	return 3 * z * z
}

func NewtonIter(f, fDx Unary, start complex128, nit int) complex128 {
	//oss := start
	for i := 0; i < nit; i++ {
		start = start - f(start)/fDx(start)
	}
	//log.Println("way", Dst(oss, start))
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
	var roots []complex128
	for i := 0; i < 3; i++ {
		roots = append(roots, cmplx.Rect(100, 2*math.Pi/3.*float64(i)))
	}
	img := [H][W]float64{}
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			p := NewtonIter(Hyper, HyperDx, complex(float64(x-W/2), float64(y-H/2)), 100)
			closestRoot := ClosetPoint(p, roots)
			img[y][x] = 255. / float64(len(roots)-closestRoot)
		}
	}

	drawAndSave(img, roots)
}

func drawAndSave(m [H][W]float64, roots []complex128) {
	im := image.NewRGBA(image.Rect(0, 0, W, H))
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
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
		x := int(real(r)) + W/2
		y := int(imag(r)) + H/2
		draw.Draw(im, image.Rect(x, y, x+10, y+10), image.White, image.Pt(0, 0), draw.Src)
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
