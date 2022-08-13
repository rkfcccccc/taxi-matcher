package qtree

import (
	"testing"
)

func TestQuadTree(t *testing.T) {
	tree := NewQuadTree(10, 10)

	tree.InsertPoint(&Point{0, 0, 0})
	tree.InsertPoint(&Point{1, 5, 0})
	tree.InsertPoint(&Point{2, 6, 6})
	tree.InsertPoint(&Point{3, 9, 9})
	tree.InsertPoint(&Point{4, 0, 9})

	if tree.size != 5 {
		t.Fatalf("tree has invalid size: %d", tree.size)
	}

	result := tree.QueryClosestInRange(&BoundingBox{4, 4, 7, 7}, &Point{-1, 5, 5})
	if result == nil || result.Key != 2 {
		t.Fatalf("closest point was not found: %+v", result)
	}

	tree.DeletePoint(&Point{1, 5, 0})
	tree.DeletePoint(&Point{2, 6, 6})

	if tree.size != 3 {
		t.Fatalf("tree has invalid size: %d", tree.size)
	}

	result = tree.QueryClosestInRange(&BoundingBox{0, 0, 10, 10}, &Point{-1, 5, 5})
	if result == nil || result.Key != 3 {
		t.Fatalf("closest point was not found: %+v", result)
	}

	result = tree.QueryClosestInRange(&BoundingBox{0, 0, 1, 1}, &Point{-1, 5, 5})
	if result == nil || result.Key != 0 {
		t.Fatalf("closest point was not found: %+v", result)
	}

	tree.DeletePoint(&Point{0, 0, 0})

	result = tree.QueryClosestInRange(&BoundingBox{0, 0, 1, 1}, &Point{-1, 0, 0})
	if result != nil {
		t.Fatalf("closest point was found, but it should not: %+v", result)
	}

	tree.DeletePoint(&Point{3, 9, 9})
	tree.DeletePoint(&Point{4, 0, 9})

	if tree.size != 0 {
		t.Fatalf("size should be 0 but it is %d", tree.size)
	}

	for i := 0; i < 20; i++ {
		tree.InsertPoint(&Point{i, 4, 4})
	}
}

func TestExportPoints(t *testing.T) {
	N := 100

	treeA := NewQuadTree(10, 10)
	for i := 0; i < N; i++ {
		treeA.InsertPoint(&Point{i, (i * i) % 10, (i * i * i) % 10})
	}

	treeB := NewQuadTree(10, 10)
	treeA.exportPoints(treeB.points)

	if len(treeB.points) != N {
		t.Fatalf("expected points count to be 20 but got %d", len(treeB.points))
	}
}
