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
	f := perlin.Field(W, H)
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
