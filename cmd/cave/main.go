package main

import (
	"math/rand"
	"newton-pcg/cave"
	"newton-pcg/core"
)

const (
	H = 240 * 2
	W = 720 * 2
)

func main() {
	rand.Seed(42)
	//n := 6
	img1 := core.RandomPool(6, W, H, core.GenerationOpts{
		Scale:     15,
		A:         complex(4, 2),
		Offset:    complex(-300, -100),
		Nit:       8,
		Metric:    &core.Mankhaten,
		DstThresh: core.Ptr(3.),
	})
	img2 := core.RandomPool(3, W, H, core.GenerationOpts{
		Scale:     50,
		A:         complex(2, 2),
		Offset:    complex(50, 200),
		Nit:       20,
		DstThresh: core.Ptr(10.),
	})

	oreMask := cave.SurfaceMask(W, H, 220)
	img2 = core.Mul(img2, oreMask)
	img2 = core.AddInt(img2, 6)

	img1 = core.AddInt(img1, 1)
	img1 = core.Overlay(img1, img2, 6)

	mask := cave.SurfaceMask(W, H, 120)
	img1 = core.Mul(img1, mask)

	//cavePool := core.RandomPool(5, W, H, core.GenerationOpts{
	//	Scale:  20,
	//	A:      complex(2, 1),
	//	Offset: complex(0, 5),
	//	Nit:    10,
	//})
	//caveMask := cave.CavesMask(cavePool, 2)
	//img1 = core.Mul(img1, caveMask)

	cave.DrawWorld(img1)
}
