package core

import "math"

type Metric = func(complex128, complex128) float64

func Euclyd(p1, p2 complex128) float64 {
	dx := real(p1 - p2)
	dy := imag(p1 - p2)
	return math.Sqrt(dx*dx + dy*dy)
}

var Mankhaten = func(p1, p2 complex128) float64 {
	dx := real(p1 - p2)
	dy := imag(p1 - p2)
	return math.Abs(dx) + math.Abs(dy)
}

var LInf = func(p1, p2 complex128) float64 {
	return math.Abs(real(p1-p2)) + math.Abs(imag(p1-p2))
}

func ThreashMetricClosetPoint(p complex128, ps []complex128, Dst Metric, thr float64) int {
	cls := 0
	dst := Dst(p, ps[0])
	for i, p2 := range ps {
		d := Dst(p, p2)
		if d < dst && d < thr {
			dst = d
			cls = i
		}
	}
	return cls
}
