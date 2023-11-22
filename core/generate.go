package core

import "math"

type Unary func(complex128) complex128

type GenerationOpts struct {
	Scale     float64
	A         complex128
	Offset    complex128
	Nit       int
	Metric    *Metric
	DstThresh *float64
}

func RandomPool(n, W, H int, opts GenerationOpts) Field {
	poly1 := RandomPoly(n, opts.Scale)
	return GeneratePool(poly1, W, H, opts)
}

func GeneratePool(poly Poly, w, h int, opts GenerationOpts) Field {
	polyPrime := poly.Prime()
	roots := poly.Roots()
	img := make([][]int, h, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w, w)
	}

	MeshCB4G(w, h, func(p P) {
		x, y := p.X, p.Y
		xx := (float64(x) + real(opts.Offset)) / opts.Scale
		yy := (float64(y) + imag(opts.Offset)) / opts.Scale
		pp := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), opts.Nit, opts.A)

		var metric = Euclyd
		if opts.Metric != nil {
			metric = *opts.Metric
		}
		var thr = math.MaxFloat64
		if opts.DstThresh != nil {
			thr = *opts.DstThresh
		}

		closestRoot := ThreashMetricClosetPoint(pp, roots, metric, thr)
		img[y][x] = closestRoot
	})
	return Field{
		W: w,
		H: h,
		F: img,
	}
}

func GenerateSpeedPool(poly Poly, w, h int, opts GenerationOpts) FField {
	polyPrime := poly.Prime()
	roots := poly.Roots()
	img := make([][]float64, h, h)
	for i := 0; i < h; i++ {
		img[i] = make([]float64, w, w)
	}
	var mx float64
	var s float64
	for p := range Mesh(w, h) {
		x, y := p.X, p.Y
		xx := (float64(x) + real(opts.Offset)) / opts.Scale
		yy := (float64(y) + imag(opts.Offset)) / opts.Scale
		p1 := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), opts.Nit, opts.A)

		var metric = Euclyd
		if opts.Metric != nil {
			metric = *opts.Metric
		}
		var thr = math.MaxFloat64
		if opts.DstThresh != nil {
			thr = *opts.DstThresh
		}
		cls := ThreashMetricClosetPoint(p1, roots, metric, thr)
		v := metric(roots[cls], p1)
		v = 1 / (1 + v)
		//v = math.Exp(-(v * v))
		//v = SmoothStep(v, 0, 10)
		img[y][x] = v
		if v > mx {
			mx = v
		}
		s += v

		//for r := range roots {
		//	img[y][x] += int(metric(roots[r], p))
		//}
		//closestRoot := ThreashMetricClosetPoint(p, roots, metric, thr)
	}
	for p := range Mesh(w, h) {
		v := img[p.Y][p.X]
		v = v - s/float64(w*h)
		img[p.Y][p.X] = SmoothStep(math.Exp(-(v * v)), 1, 2)
	}
	return FField{
		W: w,
		H: h,
		F: img,
	}
	//println("c", c1, "v", c2)
	//return Field{
	//	W: w,
	//	H: h,
	//	F: img,
	//}
}

func SmoothStep(t, start, end float64) float64 {
	t = math.Pow(t, 2) * (3 - (2 * t))
	return Linear(t, start, end)
}

func Linear(t, start, end float64) float64 {
	return t*(end-start) + start
}
