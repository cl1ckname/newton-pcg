package main

import (
	"newton-pcg/core"
	"newton-pcg/pic"
)

func main() {
	n := 3
	p := core.CirclePoly(n, 100)
	img1 := core.GeneratePool(p, 800, 800, core.GenerationOpts{
		Scale:  1,
		A:      complex(2, 1),
		Offset: complex(-400, -400),
		Nit:    10,
	})
	//img2 := core.RandomPool(n, 800, 800, core.GenerationOpts{
	//	Scale:  10,
	//	A:      1,
	//	Offset: 0,
	//	Nit:    15,
	//})
	//img3 := core.RandomPool(n, 800, 800, core.GenerationOpts{
	//	Scale:  5,
	//	A:      complex(2, 1),
	//	Offset: 0,
	//	Nit:    15,
	//})
	//img1 = core.Mix(img1, img2, n)
	//img1 = core.Mix(img1, img3, n)
	pic.DrawAndSave(img1, []complex128{})
}
