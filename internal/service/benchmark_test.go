package service

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	cases := []struct{ drivers, dim int }{
		{1e2, 1e5}, {1e3, 1e5}, {1e1, 1e6}, {1e6, 1e3},
	}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d_drivers_%d_mapdimension", c.drivers, c.dim), func(b *testing.B) {
			b.StopTimer()
			service := NewService(c.dim)

			for i := 0; i < c.drivers; i++ {
				service.Add(&Driver{Id: i, X: rand.Intn(c.dim), Y: rand.Intn(c.dim)})
			}

			queries := make([]*Pedestrian, b.N)
			for i := 0; i < b.N; i++ {
				queries[i] = &Pedestrian{X: rand.Intn(c.dim), Y: rand.Intn(c.dim)}
			}

			b.StartTimer()

			for i := 0; i < b.N; i++ {
				service.SearchClosest(queries[i])
			}
		})
	}

}
