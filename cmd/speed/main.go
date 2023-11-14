package main

import (
	"math/rand"
	"newton-pcg/core"
	"newton-pcg/speed"
	"os"
	"strconv"
)

const (
	W = 720
	H = 720
	N = 15
)

func main() {
	rand.Seed(42)
	poly1 := core.RandomPoly(N, 15)
	sp1 := core.GenerateSpeedPool(poly1, W, H, core.GenerationOpts{
		Scale:  7000,
		A:      complex(4, 4),
		Offset: complex(-360, -380),
		Nit:    11,
		//Metric: core.Ptr(core.Mankhaten),
	})

	//sp2 := core.GenerateSpeedPool(poly1, W, H, core.GenerationOpts{
	//	Scale:     1000,
	//	A:         complex(1, 4),
	//	Offset:    complex(-30, -380),
	//	Nit:       8,
	//	Metric:    nil,
	//	DstThresh: nil,
	//})
	//sp1 = core.Sum(sp1, sp2)

	//for p := range core.Mesh(W, H) {
	//	if sp1[p.Y][p.X] < 15 || sp1[p.Y][p.X] > 20 {
	//		sp1[p.Y][p.X] = 0
	//	}
	//}

	for p := range core.Mesh(W, H) {
		if d := sp1[p.Y][p.X]; d > 30 {
			sp1[p.Y][p.X] = 30
		}
	}

	f, err := os.Create("matrix.txt")
	if err != nil {
		panic(err)
	}
	for _, row := range sp1 {
		for _, col := range row {
			f.Write([]byte(strconv.FormatInt(int64(col), 10)))
			f.Write([]byte(", "))
		}
		f.Write([]byte("\n"))
	}
	if err := f.Close(); err != nil {
		panic(err)
	}

	speed.DrawAndSave(sp1)
}
