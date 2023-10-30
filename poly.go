package npcg

import (
	"math/rand"
)

type Poly struct {
	comps []float64
	roots []complex128
}

func mypow(z complex128, y int) (res complex128) {
	res = complex(1, 0)
	for i := 0; i < y; i++ {
		res *= z
	}
	return
}

func (p Poly) Eval(z complex128) complex128 {
	var res complex128
	for i, a := range p.comps {
		c := mypow(z, i)
		res += complex(a, 0) * c
	}
	return res
}

func (p Poly) Prime() Poly {
	newComps := make([]float64, len(p.comps)-1)
	for i := 0; i <= len(p.comps)-2; i++ {
		newComps[i] = p.comps[i+1] * float64(i+1)
	}
	return Poly{newComps, nil}
}

func (p Poly) N() int {
	return len(p.comps) - 1
}

func (p Poly) Roots() []complex128 {
	return p.roots
}

func RandomPoly(n int, scale float64) Poly {
	roots := make([]complex128, 0, n)
	i := n
	if i%2 == 1 {
		realRoot := complex(rand.Float64()*scale, 0)
		roots = append(roots, realRoot)
		i--
	}
	for i > 0 {
		re := rand.Float64() * scale
		im := rand.Float64() * scale
		x1 := complex(re, im)
		x2 := complex(re, -im)
		roots = append(roots, x1, x2)
		i -= 2
	}
	coefs := compsFromRoots(roots)
	return Poly{coefs, roots}
}

func FromRoots(roots []complex128) Poly {
	return Poly{
		comps: compsFromRoots(roots),
		roots: roots,
	}
}

func compsFromRoots(roots []complex128) []float64 {
	coefs := make([]complex128, len(roots)+1)
	coefs[1] = complex(1, 0)
	coefs[0] = -roots[0]
	for i := 1; i < len(roots); i++ {
		r := roots[i]
		for j := i; j >= 0; j-- {
			coefs[j+1] = coefs[j]
		}
		coefs[0] = 0
		for j := 0; j <= i; j++ {
			coefs[j] += coefs[j+1] * (-r)
		}
	}
	compsR := make([]float64, len(coefs))
	for i, c := range coefs {
		compsR[i] = real(c)
	}
	return compsR
}
