package npcg

import "math"

func NewtonIter(f, fDx Unary, start complex128, nit int, a complex128) complex128 {
	for i := 0; i < nit; i++ {
		start = start - a*f(start)/fDx(start)
	}
	return start
}

func Dst(p1, p2 complex128) float64 {
	dx := real(p1) - real(p2)
	dy := imag(p1) - imag(p2)
	return math.Sqrt(dx*dx + dy*dy)
}

func ClosetPoint(p complex128, ps []complex128) int {
	cls := 0
	dst := Dst(p, ps[0])
	for i, p2 := range ps {
		d := Dst(p, p2)
		if d < dst {
			dst = d
			cls = i
		}
	}
	return cls
}

type P struct {
	X, Y int
}

func mesh(w, h int) <-chan P {
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
	max := uint32(h.V) * 255
	min := uint32(h.V) * uint32(255-h.S)

	h.H %= 360
	segment := h.H / 60
	offset := uint32(h.H % 60)
	mid := ((max - min) * offset) / 60

	switch segment {
	case 0:
		return max, min + mid, min, 0xffff
	case 1:
		return max - mid, max, min, 0xffff
	case 2:
		return min, max, min + mid, 0xffff
	case 3:
		return min, max - mid, max, 0xffff
	case 4:
		return min + mid, min, max, 0xffff
	case 5:
		return max, min, max - mid, 0xffff
	}

	return 0, 0, 0, 0xffff
}

func Mix(img1, img2 [][]int, n int) [][]int {
	H := len(img1)
	W := len(img1[0])
	for p := range mesh(W, H) {
		img1[p.Y][p.X] = (img1[p.Y][p.X] + img2[p.Y][p.X]) % n
	}
	return img1
}
