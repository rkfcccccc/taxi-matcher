package qtree

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func closest(center *Point, a *Point, b *Point) *Point {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if center.distance2(a) < center.distance2(b) {
		return a
	} else {
		return b
	}
}
