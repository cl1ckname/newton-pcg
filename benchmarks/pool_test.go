package benchmarks

import (
	"fmt"
	"io"
	"newton-pcg/core"
	"testing"
)

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := core.RandomPool(5, W, H, core.GenerationOpts{
			Scale:     100,
			A:         complex(2, 1),
			Offset:    0,
			Nit:       5,
			Metric:    nil,
			DstThresh: nil,
		})
		_, err := fmt.Fprintf(io.Discard, "%v", v)
		if err != nil {
			panic(err)
		}
	}
}
