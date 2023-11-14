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

	for p := range Mesh(w, h) {
		x, y := p.X, p.Y
		xx := (float64(x) + real(opts.Offset)) / opts.Scale
		yy := (float64(y) + imag(opts.Offset)) / opts.Scale
		p := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), opts.Nit, opts.A)

		var metric = Euclyd
		if opts.Metric != nil {
			metric = *opts.Metric
		}
		var thr = math.MaxFloat64
		if opts.DstThresh != nil {
			thr = *opts.DstThresh
		}

		closestRoot := ThreashMetricClosetPoint(p, roots, metric, thr)
		img[y][x] = closestRoot
	}
	return Field{
		W: w,
		H: h,
		F: img,
	}
}

func GenerateSpeedPool(poly Poly, w, h int, opts GenerationOpts) Field {
	polyPrime := poly.Prime()
	roots := poly.Roots()
	img := make([][]int, h, h)
	for i := 0; i < h; i++ {
		img[i] = make([]int, w, w)
	}

	for p := range Mesh(w, h) {
		x, y := p.X, p.Y
		xx := (float64(x) + real(opts.Offset)) / opts.Scale
		yy := (float64(y) + imag(opts.Offset)) / opts.Scale
		p := NewtonIter(poly.Eval, polyPrime.Eval, complex(xx, yy), opts.Nit, opts.A)

		var metric = Euclyd
		if opts.Metric != nil {
			metric = *opts.Metric
		}
		//var thr = math.MaxFloat64
		//if opts.DstThresh != nil {
		//	thr = *opts.DstThresh
		//}
		//img[y][x] = int(metric(complex(xx, yy), p))
		for r := range roots {
			img[y][x] += int(metric(roots[r], p))
		}
		img[y][x] = img[y][x] / len(roots)
		//closestRoot := ThreashMetricClosetPoint(p, roots, metric, thr)
	}
	return Field{
		W: w,
		H: h,
		F: img,
	}
}
