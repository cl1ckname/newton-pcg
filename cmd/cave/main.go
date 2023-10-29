package main

import np "newton-pcg"

const (
	H = 320
	W = 480
)

func main() {
	n := 6
	img1 := np.RandomPool(n, W, H, 1)
	img2 := np.RandomPool(n, W, H, 1)
	img1 = np.Mix(img1, img2, n)
	img3 := np.RandomPool(n, W, H, 1)
	img1 = np.Mix(img1, img3, n)
	np.GenerateCave(img1)
}
