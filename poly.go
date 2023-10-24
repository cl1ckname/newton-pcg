package main

import (
	"math"
	"math/cmplx"
	"math/rand"
)

type Poly struct {
	comps []float64
}

func (p Poly) Eval(z complex128) complex128 {
	var res complex128
	for i, a := range p.comps {
		c := cmplx.Pow(z, complex(float64(i), 0))
		res += complex(a, 0) * c
	}
	return res
}

func (p Poly) Prime() Poly {
	newComps := make([]float64, len(p.comps)-1)
	for i := 0; i <= len(p.comps)-1; i++ {
		newComps[i] = p.comps[i+1] * float64(i+1)
	}
	return Poly{newComps}
}

func (p Poly) N() int {
	return len(p.comps) + 1
}

func (p Poly) Roots() []complex128 {
	return bairstow(p, 1, 1)
}

func bairstow(poly Poly, r, s float64) (roots []complex128) {
	if poly.N() < 1 {
		return nil
	}
	if poly.N() == 1 && poly.comps[1] != 0 {
		roots = append(roots, complex(-poly.comps[0]/poly.comps[1], 0))
		return
	}
	if poly.N() == 2 {
		c, b, a := poly.comps[0], poly.comps[1], poly.comps[2]
		D := b*b - 4*a*c
		x1 := (-complex(b, 0) - cmplx.Sqrt(complex(D, 0))) / (2 * complex(a, 0))
		x2 := (-complex(b, 0) + cmplx.Sqrt(complex(D, 0))) / (2 * complex(a, 0))
		roots = append(roots, x1, x2)
		return
	}
	a := poly.comps
	n := len(poly.comps)
	b := make([]float64, n)
	b[n-1] = a[n-1]
	b[n-2] = a[n-2] + r*b[n-1]
	i := n - 3
	for i >= 0 {
		b[i] = a[i] + r*b[i+1] + s*b[i+2]
		i--
	}
	c := make([]float64, n)
	c[n-1] = b[n-1]
	c[n-2] = b[n-2] + r*c[n-1]
	i = n - 3
	for i >= 0 {
		c[i] = b[i] + r*c[i+1] + s*c[i+2]
		i--
	}
	din := 1 / (c[2]*c[2] - c[3]*c[1])
	r = r - din*c[2]*b[1] + c[3]*b[0]
	s = s + din*c[1]*b[1] - c[2]*b[0]
	if math.Abs(b[0]) > 1e-14 || math.Abs(b[1]) > 1e-14 {
		return bairstow(poly, r, s)
	}
	if poly.N() >= 3 {
		dis := r*r + 4*s
		x1 := (complex(r, 0) - cmplx.Sqrt(complex(dis, 0))) / 2
		x2 := (complex(r, 0) + cmplx.Sqrt(complex(dis, 0))) / 2
		roots = append(roots, x1, x2)
		reducedP := Poly{b[2:]}
		reducedRoots := bairstow(reducedP, r, s)
		roots = append(roots, reducedRoots...)
	}
	return
}

func RandomPoly() Poly {
	n := rand.Intn(8)
	coefs := make([]float64, n, n)
	for i := 0; i < n; i++ {
		coef := rand.Float64() * 10
		coefs = append(coefs, coef)
	}
	return Poly{coefs}
}
