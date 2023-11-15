package main

import (
	"image"
	"math"
	"newton-pcg/core"
	"newton-pcg/perlin"
)

const (
	W = 480
	H = 480
)

func main() {
	f := core.NewField(W, H)

	perl := perlin.NewPerlin(1, 4, 3, 5)
	for p := range core.Mesh(W, H) {
		pv := int(perl.Noise2D(float64(p.X)/W, float64(p.Y)/H) * 255)
		var v int
		if pv < 120 && pv > 60 {
			v = 255
		}
		f.Set(p, v)
	}

	//b, _ := json.Marshal(f)
	//println(string(b))
	DrawAndSave(f)
}

func DrawAndSave(m core.Field) {
	im := image.NewRGBA(image.Rect(0, 0, m.W, m.H))
	for p := range core.Mesh(m.W, m.H) {
		im.Set(p.X, p.Y, core.HSVColor{
			H: math.MaxUint16,
			S: uint8(m.At(p)),
			V: uint8(m.At(p)),
		})
	}
	core.SaveImage(im)
}
