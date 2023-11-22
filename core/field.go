package core

func NewField(w, h int) Field {
	var m [][]int
	for i := 0; i < h; i++ {
		m = append(m, make([]int, w))
	}
	return Field{
		W: w,
		H: h,
		F: m,
	}
}

func NewFField(w, h int) FField {
	var m [][]float64
	for i := 0; i < h; i++ {
		m = append(m, make([]float64, w))
	}
	return FField{
		W: w,
		H: h,
		F: m,
	}
}

type Field struct {
	W int
	H int
	F [][]int
}

func (f *Field) At(p P) int {
	return f.F[p.Y][p.X]
}

func (f *Field) Set(at P, v int) {
	f.F[at.Y][at.X] = v
}

type FField struct {
	W int
	H int
	F [][]float64
}

func (f *FField) At(p P) float64 {
	return f.F[p.Y][p.X]
}

func (f *FField) Set(at P, v float64) {
	f.F[at.Y][at.X] = v
}
