package voronoi

import (
	"math"
	"math/rand"
	"newton-pcg/core"
)

func Random(w, h, n int) core.Field {
	points := make([]core.P, n)
	for i := 0; i < n; i++ {
		points[i] = core.P{
			X: rand.Intn(w),
			Y: rand.Intn(h),
		}
	}
	return FromPoints(w, h, points)
}

func FromPoints(w, h int, points []core.P) core.Field {
	f := core.NewField(w, h)

	core.MeshCB4G(w, h, func(p core.P) {
		m := math.MaxFloat64
		var mi int
		for i, pi := range points {
			if dst := core.Euclyd(complex(float64(p.X), float64(p.Y)), complex(float64(pi.X), float64(pi.Y))); dst < m {
				m = dst
				mi = i
			}
		}
		f.Set(p, mi)
	})
	return f
}
