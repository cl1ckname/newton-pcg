package pic

import (
	"image"
	"image/color"
	"image/color/palette"
	"math"
	"newton-pcg/core"
)

func DrawAndSave(m core.Field, roots []complex128) {
	im := image.NewRGBA(image.Rect(0, 0, m.W, m.H))
	for p := range core.Mesh(m.W, m.H) {
		//a := uint16(math.MaxUint8 / float64(len(roots)-m.At(p)))
		a := uint8(math.MaxUint8 / (1 + m.At(p)))
		im.Set(p.X, p.Y, color.RGBA{
			R: a,
			G: a,
			B: a,
			A: 255,
		})
	}

	for _, r := range roots {
		x := int(real(r))
		y := int(imag(r))

		for i := x; i < x+10; i++ {
			for j := y; j < y+10; j++ {
				im.Set(i, j, color.RGBA{R: 200, A: 255})
			}
		}
	}
	core.SaveImage(im)
}

func ToImage(m core.Field, roots []complex128) *image.Paletted {
	im := image.NewPaletted(image.Rect(0, 0, m.W, m.H), palette.Plan9[:256])
	for p := range core.Mesh(m.W, m.H) {
		a := uint8(255/float64(len(roots))) * uint8(m.At(p))
		im.Set(p.X, p.Y, core.HSVColor{
			H: uint16(a),
			S: 128,
			V: 224,
		})
	}
	for _, r := range roots {
		x := int(real(r)) + m.W/2
		y := int(imag(r)) + m.H/2

		for i := x; i < x+10; i++ {
			for j := y; j < y+10; j++ {
				im.Set(i, j, color.RGBA{R: 200, A: 255})
			}
		}
	}
	return im
}
