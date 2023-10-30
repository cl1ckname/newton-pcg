package main

import (
	"image"
	"image/gif"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	np "newton-pcg"
	"os"
	"sync"
)

func rootRotation(c, start complex128) (res [frames]complex128) {
	phst := 2 * math.Pi / frames
	if rand.Float32() > 0.5 {
		phst *= -1
	}
	rad := np.Dst(c, start)
	for i := 0; i < frames; i++ {
		u := cmplx.Rect(rad, float64(i)*phst)
		res[i] = u + c
	}
	return
}

const (
	n      = 15
	w      = 720
	h      = 720
	frames = 120
)

func main() {
	p := np.RandomPoly(n, w)
	roots := p.Roots()
	rotations := [n][frames]complex128{}
	for i := 0; i < n; i++ {
		r := 30 + rand.Float64()*160
		start := roots[i]
		rotations[i] = rootRotation(start-complex(0, r), start)
	}

	var fs [frames]*image.Paletted
	delays := make([]int, frames)

	var wg sync.WaitGroup
	wg.Add(frames)

	for frame := 0; frame < frames; frame++ {
		go func(frame int) {
			var newRoots [n]complex128
			for j := 0; j < n; j++ {
				newRoots[j] = rotations[j][frame]
			}
			poly := np.FromRoots(newRoots[:])
			pool := np.GeneratePool(poly, w, h, np.GenerationOpts{
				Scale:  1,
				A:      complex(4, 2),
				Offset: 0,
				Nit:    3,
			})
			img := np.ToImage(pool, newRoots[:])
			fs[frame] = img
			wg.Done()
		}(frame)
	}
	wg.Wait()
	f, err := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := gif.EncodeAll(f, &gif.GIF{
		Image: fs[:],
		Delay: delays,
	}); err != nil {
		log.Fatal(err)
	}
}
