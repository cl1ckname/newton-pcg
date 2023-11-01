package main

import (
	np "newton-pcg/cave"
	np2 "newton-pcg/np"
)

const (
	H = 160
	W = 240
)

func main() {
	n := 6
	img1 := np2.RandomPool(n, W, H, np2.GenerationOpts{
		Scale:  10,
		A:      complex(2, 1),
		Offset: complex(-4, 3),
		Nit:    5,
	})
	//img2 := np.RandomPool(n, W, H, np.GenerationOpts{
	//	Scale:  15,
	//	A:      complex(2, 1),
	//	Offset: complex(5, -4),
	//	Nit:    15,
	//})
	//img1 = np.Mix(img1, img2, n)
	//img3 := np.RandomPool(n, W, H, np.GenerationOpts{
	//	Scale:  5,
	//	A:      complex(2, 1),
	//	Offset: complex(5, 5),
	//	Nit:    20,
	//})
	//img1 = np.Mix(img1, img3, n)
	np.GenerateCave(img1)
}
