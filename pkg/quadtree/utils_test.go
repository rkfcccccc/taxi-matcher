package qtree

import "testing"

func TestMin(t *testing.T) {
	if min(1, 2) != 1 {
		t.Fatalf("min(1, 2) != 1")
	}

	if min(2, 1) != 1 {
		t.Fatalf("min(2, 1) != 1")
	}
}

func TestMax(t *testing.T) {
	if max(1, 2) != 2 {
		t.Fatalf("max(1, 2) != 2")
	}

	if max(2, 1) != 2 {
		t.Fatalf("max(2, 1) != 2")
	}
}

func TestClosest(t *testing.T) {
	center := &Point{X: 1, Y: 1}
	a := &Point{X: 1, Y: 1}
	b := &Point{X: -100, Y: -100}

	if closest(center, a, b) != a {
		t.Fatalf("closest(center, a, b) != a")
	}

	if closest(center, b, a) != a {
		t.Fatalf("closest(center, b, a) != a")
	}

	if closest(center, center, nil) != center {
		t.Fatalf("closest(center, center, nil) != center")
	}

	if closest(center, nil, center) != center {
		t.Fatalf("closest(center, nil, center) != center")
	}
}
