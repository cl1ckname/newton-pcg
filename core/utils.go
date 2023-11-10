package core

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func NewtonIter(f, fDx Unary, start complex128, nit int, a complex128) complex128 {
	for i := 0; i < nit; i++ {
		start = start - a*f(start)/fDx(start)
	}
	return start
}

type P struct {
	X, Y int
}

func Mesh(w, h int) <-chan P {
	c := make(chan P, w*h)
	go func() {
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				c <- P{x, y}
			}
		}
		close(c)
	}()
	return c
}

type HSVColor struct {
	H    uint16
	S, V uint8
}

func (h HSVColor) RGBA() (r, g, b, a uint32) {
	// Direct implementation of the graph in this image:
	// https://en.wikipedia.org/wiki/HSL_and_HSV#/media/File:HSV-RGB-comparison.svg
	mx := uint32(h.V) * 255
	mn := uint32(h.V) * uint32(255-h.S)

	h.H %= 360
	segment := h.H / 60
	offset := uint32(h.H % 60)
	mid := ((mx - mn) * offset) / 60

	switch segment {
	case 0:
		return mx, mn + mid, mn, 0xffff
	case 1:
		return mx - mid, mx, mn, 0xffff
	case 2:
		return mn, mx, mn + mid, 0xffff
	case 3:
		return mn, mx - mid, mx, 0xffff
	case 4:
		return mn + mid, mn, mx, 0xffff
	case 5:
		return mx, mn, mx - mid, 0xffff
	}

	return 0, 0, 0, 0xffff
}

func Mix(img1, img2 [][]int, n int) [][]int {
	H := len(img1)
	W := len(img1[0])
	for p := range Mesh(W, H) {
		img1[p.Y][p.X] = (img1[p.Y][p.X] + img2[p.Y][p.X]) % n
	}
	return img1
}

func Mul(img1, img2 [][]int) [][]int {
	H := len(img1)
	W := len(img1[0])
	for p := range Mesh(W, H) {
		img1[p.Y][p.X] *= img2[p.Y][p.X]
	}
	return img1
}

func Sum(img1, img2 [][]int) [][]int {
	H := len(img1)
	W := len(img1[0])
	for p := range Mesh(W, H) {
		img1[p.Y][p.X] += img2[p.Y][p.X]
	}
	return img1
}

func Overlay(img1, img2 [][]int, b int) [][]int {
	H := len(img1)
	W := len(img1[0])
	for p := range Mesh(W, H) {
		v := img2[p.Y][p.X]
		if v != b {
			img1[p.Y][p.X] = v
		}
	}
	return img1
}

func AddInt(img1 [][]int, i int) [][]int {
	H := len(img1)
	W := len(img1[0])
	for p := range Mesh(W, H) {
		img1[p.Y][p.X] += i
	}
	return img1
}

func SaveImage(img image.Image) {
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

func Ptr[T any](v T) *T {
	return &v
}
