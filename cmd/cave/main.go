package main

import (
	"newton-pcg/cave"
	"newton-pcg/core"
)

const (
	H = 240 * 2
	W = 720 * 2
)

func main() {
	n := 6
	img1 := core.RandomPool(n, W, H, core.GenerationOpts{
		Scale:  50,
		A:      complex(4, 2),
		Offset: complex(-4, 4),
		Nit:    5,
	})
	img2 := core.RandomPool(n, W, H, core.GenerationOpts{
		Scale:  15,
		A:      complex(2, 1),
		Offset: complex(5, -4),
		Nit:    15,
	})
	//img1 = np.Mix(img1, img2, n)
	//img3 := np.RandomPool(n, W, H, np.GenerationOpts{
	//	Scale:  5,
	//	A:      complex(2, 1),
	//	Offset: complex(5, 5),
	//	Nit:    20,
	//})
	img1 = core.Mix(img1, img2, n)
	img1 = core.AddInt(img1, 1)
	mask := cave.SurfaceMask(W, H, 120)
	img1 = core.Mul(img1, mask)

	cavePool := core.RandomPool(5, W, H, core.GenerationOpts{
		Scale:  20,
		A:      complex(2, 1),
		Offset: complex(0, 5),
		Nit:    10,
	})
	caveMask := cave.CavesMask(cavePool, 2)
	img1 = core.Mul(img1, caveMask)

	cave.DrawWorld(img1)
}
