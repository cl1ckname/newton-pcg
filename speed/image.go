package speed

import (
	"image"
	"math"
	"newton-pcg/core"
)

func DrawAndSave(m [][]int) {
	h := len(m)
	w := len(m[0])

	im := image.NewRGBA(image.Rect(0, 0, w, h))

	var mx float64
	for p := range core.Mesh(w, h) {
		if m := float64(m[p.Y][p.X]); m > mx {
			mx = m
		}
	}
	//println(mx)
	norm := math.MaxUint16 / mx

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a := float64(m[x][y]) * norm
			im.Set(x, y, core.HSVColor{
				H: uint16(a),
				S: 255,
				V: 255,
			})
		}
	}
	core.SaveImage(im)
}