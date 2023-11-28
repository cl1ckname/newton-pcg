package main

import (
	"newton-pcg/core"
)

func Generate(w, h int) core.Field {
	ply := core.CirclePoly(3, 5)
	img := core.GeneratePool(ply, w, h, core.GenerationOpts{
		Scale:     500,
		A:         complex(2, 1),
		Offset:    complex(-float64(w)/2, -float64(h)/2+200),
		Nit:       9,
		Metric:    core.Ptr(core.Mankhaten),
		DstThresh: nil,
	})
	core.AddInt(img, 1)

	imgCpy := findContours(img)
	img = imgCpy
	for i := 0; i < 1; i++ {
		img = erose(img)
	}

	return img
}

func findContours(img core.Field) core.Field {
	w := img.W
	h := img.H
	imgCpy := core.NewField(w, h)

	for p := range core.Mesh(w, h) {
		var l, r, b, t bool
		if p.X == 0 {
			l = true
		}
		if p.Y == 0 {
			t = true
		}
		if p.X == w-1 {
			r = true
		}
		if p.Y == h-1 {
			b = true
		}
		if l || b || r || t {
			continue
		}

		v := img.At(p.Top())
		if img.At(p.Bottom()) == v && img.At(p.Left()) == v && img.At(p.Right()) == v {
			imgCpy.Set(p, img.At(p))
			continue
		}
	}
	return imgCpy
}

func erose(img core.Field) core.Field {
	w := img.W
	h := img.H
	imgCpy := core.NewField(w, h)
	for p := range core.Mesh(w-2, h-2) {
		pn := p.Bottom().Right()
		if img.At(pn) == 0 {
			continue
		}
		if img.At(pn.Top()) == 0 {
			continue
		}
		if img.At(pn.Left()) == 0 {
			continue
		}
		if img.At(pn.Right()) == 0 {
			continue
		}
		if img.At(pn.Bottom()) == 0 {
			continue
		}
		imgCpy.Set(pn, img.At(pn))
	}
	return imgCpy
}
