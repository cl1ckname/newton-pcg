package benchmarks

import (
	"fmt"
	"io"
	"newton-pcg/voronoi"
	"testing"
)

const (
	W = 1000
	H = 1000
	N = 1000
)

func BenchmarkVoronoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := voronoi.Random(W, H, N)
		_, err := fmt.Fprintf(io.Discard, "%v", v)
		if err != nil {
			panic(err)
		}
	}
}
