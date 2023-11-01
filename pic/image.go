package pic

import (
	"image"
	"image/color"
	"image/color/palette"
	"newton-pcg/core"
)

func DrawAndSave(m [][]int, roots []complex128) {
	h := len(m)
	w := len(m[0])

	im := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a := uint8(255 / float64(len(roots)-m[x][y]))
			im.Set(x, y, core.HSVColor{
				H: uint16(a) * 256,
				S: 128,
				V: 224,
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
	core.SaveImage(im)
}

func ToImage(m [][]int, roots []complex128) *image.Paletted {
	h := len(m)
	w := len(m[0])

	im := image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9[:256])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a := uint8(255/float64(len(roots))) * uint8(m[x][y])
			im.Set(x, y, core.HSVColor{
				H: uint16(a),
				S: 128,
				V: 224,
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
	}
	return im
}