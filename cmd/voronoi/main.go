package main

import (
	"math/rand"
	"newton-pcg/core"
	"newton-pcg/pic"
	"newton-pcg/voronoi"
)

const (
	W = 720 * 2
	H = 720 * 2
	N = 1000
)

func main() {

	points := make([]core.P, N)
	for i := 0; i < N; i++ {
		points[i] = core.P{
			X: rand.Intn(W),
			Y: rand.Intn(H),
		}
	}

	//f := voronoi.Random(W, H, N)
	f := voronoi.FromPoints(W, H, points)
	cpoints := make([]complex128, N)
	for i, p := range points {
		cpoints[i] = complex(float64(p.X), float64(p.Y))
	}
	pic.DrawAndSave(f, cpoints)
}
