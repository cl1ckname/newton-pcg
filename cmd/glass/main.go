package main

import (
	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/noise"
	"image"
	"image/color"
	"image/color/palette"
	"math/rand"
	"newton-pcg/core"
)

const (
	W = 720
	H = 720
)

func main() {
	rand.Seed(42)
	img := Generate(W, H)
	lena := core.GetImageFromFilePath("assets/Lenna.png")
	ToImage(img, lena)
}

func ToImage(m core.Field, over image.Image) {
	im := image.NewPaletted(image.Rect(0, 0, m.W, m.H), palette.Plan9[:256])
	for p := range core.Mesh(m.W, m.H) {
		var c color.RGBA
		switch m.At(p) % 5 {
		case 0:
			c = color.RGBA{}
		case 1:
			c = color.RGBA{R: 224, G: 64, B: 128}
		case 2:
			c = color.RGBA{G: 224, B: 64, R: 128}
		case 3:
			c = color.RGBA{B: 224, R: 64, G: 128}
		default:
			c = color.RGBA{G: 255, B: 255, R: 255}
		}

		im.Set(p.X, p.Y, c)
	}
	img := GlassFilter(im, over)
	core.SaveImage(img)
}

func GlassFilter(img *image.Paletted, over image.Image) *image.RGBA {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	im := blur.Gaussian(over, 5)
	for i := 0; i < 15; i++ {
		im = blur.Gaussian(im, 10)
		im = blur.Box(im, 5)
	}

	n := noise.GeneratePerlin(w, h, 0.5)
	n = adjust.Contrast(n, -0.75)
	fg := blend.Darken(img, n)
	res := blend.Multiply(fg, im)
	return adjust.Brightness(res, 0.75)
	//draw.Draw(img, img.Bounds(), over, image.Point{}, draw.Over)
	//for p := range core.Mesh(w,h) {
	//	r1,g1,b1,a :=
	//	c := color.RGBA{
	//		R: over.,
	//		G: 0,
	//		B: 0,
	//		A: 0,
	//	}
	//	img.Set(p.X, p.Y)
	//}
}
