package service

import (
	"math/rand"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	b.StopTimer()

	dim := 100000
	service := NewService(dim)

	for i := 0; i < 1e4; i++ {
		service.Add(&Driver{Id: i, X: rand.Intn(dim), Y: rand.Intn(dim)})
	}

	queries := make([]*Pedestrian, b.N)
	for i := 0; i < b.N; i++ {
		queries[i] = &Pedestrian{X: rand.Intn(dim), Y: rand.Intn(dim)}
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		service.SearchClosest(queries[i])
	}
}
