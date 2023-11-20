package perlin

import "newton-pcg/core"

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
