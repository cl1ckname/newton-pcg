package main

import (
	"math/rand"
	"newton-pcg/cave"
	"newton-pcg/core"
	"testing"
)

func TestCave(t *testing.T) {
	rand.Seed(int64(SEED))
	//n := 6
	img1 := core.RandomPool(6, W, H, core.GenerationOpts{
		Scale:     15,
		A:         complex(4, 2),
		Offset:    complex(-300, -100),
		Nit:       8,
		Metric:    &core.Mankhaten,
		DstThresh: core.Ptr(3.),
	})
	img2 := core.RandomPool(3, W, H, core.GenerationOpts{
		Scale:     50,
		A:         complex(2, 2),
		Offset:    complex(50, 200),
		Nit:       20,
		DstThresh: core.Ptr(10.),
	})

	oreMask := cave.SurfaceMask(W, H, 220)
	img2 = core.Mul(img2, oreMask)
	img2 = core.AddInt(img2, 6)

	img1 = core.AddInt(img1, 1)
	img1 = core.Overlay(img1, img2, 6)

	mask := cave.SurfaceMask(W, H, 120)
	img1 = core.Mul(img1, mask)

	caveMask := cave.Mask(W, H, 1, 8, int64(SEED))
	img1 = core.Mul(img1, caveMask)
	caveMask = cave.Mask(W, H, 0.5, 4, int64(SEED))
	img1 = core.Mul(img1, caveMask)
	caveMask = cave.Mask(W, H, 0.5, 6, int64(SEED))
	img1 = core.Mul(img1, caveMask)

	groundMask := cave.GroundLayer(W, H, 210)
	groundMask = core.MulI(groundMask, 2)
	img1 = core.Replace(img1, groundMask)

	img1 = cave.GrassLayer(img1, 2, 10)

	cave.DrawWorld(img1)
}
