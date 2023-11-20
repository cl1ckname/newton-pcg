package benchmarks

import (
	"fmt"
	"io"
	"newton-pcg/perlin"
	"testing"
)

func BenchmarkPerlin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := perlin.Field(W, H)
		_, err := fmt.Fprintf(io.Discard, "%v", v)
		if err != nil {
			panic(err)
		}
	}

}
