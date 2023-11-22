package speed

import (
	"image"
	"math"
	"newton-pcg/core"
)

func DAS(m core.FField) {
	im := image.NewRGBA(image.Rect(0, 0, m.W, m.H))

	var mx float64
	for p := range core.Mesh(m.W, m.H) {
		if m := float64(m.At(p)); m > mx {
			mx = m
		}
	}

	for p := range core.Mesh(m.W, m.H) {
		a := float64(m.At(p))
		im.Set(p.X, p.Y, core.HSVColor{
			H: 0,
			S: 0,
			V: uint8(math.MaxUint8 * a),
		})
	}
	core.SaveImage(im)
}

func DrawAndSave(m core.Field) {
	im := image.NewRGBA(image.Rect(0, 0, m.W, m.H))

	var mx float64
	for p := range core.Mesh(m.W, m.H) {
		if m := float64(m.At(p)); m > mx {
			mx = m
		}
	}
	//println(mx)
	norm := math.MaxUint16 / mx

	for p := range core.Mesh(m.W, m.H) {
		a := float64(m.At(p)) * norm
		im.Set(p.X, p.Y, core.HSVColor{
			H: uint16(a),
			S: 255,
			V: 255,
		})
	}
	core.SaveImage(im)
}
