package core

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
