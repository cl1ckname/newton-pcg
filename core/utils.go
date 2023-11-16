package core

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"sync"
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

func Mix(img1, img2 Field, n int) Field {
	for p := range Mesh(img1.W, img1.H) {
		img1.Set(p, img1.At(p)+img2.At(p)%n)
	}
	return img1
}

func Mul(img1, img2 Field) Field {
	MeshCB4G(img1.W, img1.H, func(p P) {
		img1.Set(p, img1.At(p)*img2.At(p))
	})

	return img1
}

func MulI(img Field, i int) Field {
	for p := range Mesh(img.W, img.H) {
		img.Set(p, img.At(p)*i)
	}
	return img
}

func Sum(img1, img2 Field) Field {
	for p := range Mesh(img1.W, img1.H) {
		img1.Set(p, img1.At(p)+img2.At(p))
	}
	return img1
}

func Overlay(img1, img2 Field, b int) Field {
	for p := range Mesh(img1.W, img1.H) {
		v := img2.At(p)
		if v != b {
			img1.Set(p, v)
		}
	}
	return img1
}

func Replace(img1, img2 Field) Field {
	for p := range Mesh(img1.W, img1.H) {
		v1 := img1.At(p)
		v2 := img2.At(p)
		if v1 != 0 && v2 != 0 {
			img1.Set(p, v2)
		}
	}
	return img1
}

func AddInt(img1 Field, i int) Field {
	for p := range Mesh(img1.W, img1.H) {
		img1.Set(p, img1.At(p)+i)
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

//func Norm(arr [][]int, m int, d int) [][]int {
//
//	for p := range Mesh()
//}

func MeshCB4G(w, h int, cb func(P)) {
	wg := sync.WaitGroup{}
	wg.Add(4)
	halfIter := func(startX, endX, startY, endY int) {
		for x := startX; x < endX; x++ {
			for y := startY; y < endY; y++ {
				cb(P{x, y})
			}
		}
		wg.Done()
	}
	go halfIter(0, w/2, 0, h/2)
	go halfIter(w/2, w, 0, h/2)
	go halfIter(0, w/2, h/2, h)
	go halfIter(w/2, w, h/2, h)
	wg.Wait()
}
