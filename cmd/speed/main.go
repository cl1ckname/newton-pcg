package main

import (
	"newton-pcg/core"
	"newton-pcg/pic"
)

const (
	W = 720
	H = 720
	N = 5
)

func main() {
	poly1 := core.RandomPoly(N, 15)
	sp := core.GenerateSpeedPool(poly1, W, H, core.GenerationOpts{
		Scale:     60,
		A:         complex(2, 1),
		Offset:    complex(-360, -480),
		Nit:       5,
		Metric:    nil,
		DstThresh: nil,
	})

	pic.DrawAndSave(sp, []complex128{})
}
