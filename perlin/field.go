package perlin

import (
	"math"
	"newton-pcg/core"
)

func Field(W, H int) core.Field {
	f := core.NewField(W, H)

	perl := NewPerlin(1, 4, 3, 5)
	core.MeshCB4G(W, H, func(p core.P) {
		pv := int(perl.Noise2D(float64(p.X)/float64(W), float64(p.Y)/float64(H)) * 255)
		var v int
		if pv < 120 && pv > 60 {
			v = 255
		}
		f.Set(p, v)
	})
	return f
}

func New(W, H int, a, b float64) core.Field {
	f := core.NewField(W, H)

	perl := NewPerlin(a, b, 3, 5)
	core.MeshCB4G(W, H, func(p core.P) {
		pv := 128 + int(math.Abs(perl.Noise2D(float64(p.X)/float64(W), float64(p.Y)/float64(H))*127))
		f.Set(p, pv)
	})
	return f
}
