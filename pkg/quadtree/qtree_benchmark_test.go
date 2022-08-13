package qtree

import (
	"math/rand"
	"testing"
)

const fieldSize = 1e6

func BenchmarkInsertion(b *testing.B) {
	b.StopTimer()

	tree := NewQuadTree(fieldSize, fieldSize)
	points := make([]Point, b.N)

	for i := 0; i < b.N; i++ {
		points[i] = Point{Key: i, X: rand.Intn(fieldSize), Y: rand.Intn(fieldSize)}
	}

	b.StartTimer()

	for _, p := range points {
		tree.InsertPoint(&p)
	}
}

func BenchmarkDeletion(b *testing.B) {
	b.StopTimer()

	tree := NewQuadTree(fieldSize, fieldSize)
	points := make([]Point, b.N)

	for i := 0; i < b.N; i++ {
		points[i] = Point{Key: i, X: rand.Intn(fieldSize), Y: rand.Intn(fieldSize)}
	}

	for _, p := range points {
		tree.InsertPoint(&p)
	}

	b.StartTimer()

	for _, p := range points {
		tree.DeletePoint(&p)
	}
}
