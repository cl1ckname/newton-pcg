package npcg

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func drawAndSave(m [][]int, roots []complex128) {
	h := len(m)
	w := len(m[0])

	im := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a := uint8(255 / float64(len(roots)-m[x][y]))
			im.Set(x, y, HSVColor{
				H: uint16(a) * 256,
				S: 128,
				V: 224,
			})
		}
	}
	for _, r := range roots {
		x := int(real(r))
		y := int(imag(r))

		for i := x; i < x+10; i++ {
			for j := y; j < y+10; j++ {
				im.Set(i, j, color.RGBA{200, 0, 0, 255})
			}
		}
	}
	saveImage(im)
}

func saveImage(img image.Image) {
	f, err := os.Create("out.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	opt := jpeg.Options{Quality: 50}
	err = jpeg.Encode(f, img, &opt)
	if err != nil {
		log.Fatal(err)
	}
}
