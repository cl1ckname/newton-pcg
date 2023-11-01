package main

import (
	"newton-pcg/core"
	"newton-pcg/pic"
)

func main() {
	n := 7
	img1 := core.RandomPool(n, 800, 800, core.GenerationOpts{
		Scale:  20,
		A:      1,
		Offset: 0,
		Nit:    15,
	})
	img2 := core.RandomPool(n, 800, 800, core.GenerationOpts{
		Scale:  10,
		A:      1,
		Offset: 0,
		Nit:    15,
	})
	img3 := core.RandomPool(n, 800, 800, core.GenerationOpts{
		Scale:  5,
		A:      complex(2, 1),
		Offset: 0,
		Nit:    15,
	})
	img1 = core.Mix(img1, img2, n)
	img1 = core.Mix(img1, img3, n)
	pic.DrawAndSave(img1, []complex128{})
}
