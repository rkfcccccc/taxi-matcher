package qtree

const nodeCapacity = 4

type Point struct {
	Key  int
	X, Y int
}

func (p *Point) distance2(other *Point) int {
	return (p.X-other.X)*(p.X-other.X) + (p.Y-other.Y)*(p.Y-other.Y)
}

type BoundingBox struct {
	// [x1; x2), [y1; y2)
	X1, Y1, X2, Y2 int
}

func (bb *BoundingBox) Contains(p *Point) bool {
	return p.X >= bb.X1 && p.X < bb.X2 && p.Y >= bb.Y1 && p.Y < bb.Y2
}

func (bb *BoundingBox) Intersects(other *BoundingBox) bool {
	return max(bb.X1, other.X1) < min(bb.X2, other.X2) && max(bb.Y1, other.Y1) < min(bb.Y2, other.Y2)
}

type QuadTree struct {
	size   int
	points map[Point]struct{}

	boundingBox *BoundingBox

	northEast *QuadTree
	northWest *QuadTree
	southEast *QuadTree
	southWest *QuadTree
}

func NewQuadTree(width, height int) *QuadTree {
	return &QuadTree{boundingBox: &BoundingBox{X2: width, Y2: height}, points: map[Point]struct{}{}}
}

func (qt *QuadTree) subdivide() {
	if qt.boundingBox.X2-qt.boundingBox.X1 <= 1 {
		return
	}

	if qt.boundingBox.Y2-qt.boundingBox.Y1 <= 1 {
		return
	}

	qt.southEast = &QuadTree{boundingBox: &BoundingBox{
		X1: qt.boundingBox.X1, X2: (qt.boundingBox.X1 + qt.boundingBox.X2) / 2,
		Y1: qt.boundingBox.Y1, Y2: (qt.boundingBox.Y1 + qt.boundingBox.Y2) / 2,
	}, points: map[Point]struct{}{}}

	qt.southWest = &QuadTree{boundingBox: &BoundingBox{
		X1: (qt.boundingBox.X1 + qt.boundingBox.X2) / 2, X2: qt.boundingBox.X2,
		Y1: qt.boundingBox.Y1, Y2: (qt.boundingBox.Y1 + qt.boundingBox.Y2) / 2,
	}, points: map[Point]struct{}{}}

	qt.northEast = &QuadTree{boundingBox: &BoundingBox{
		X1: qt.boundingBox.X1, X2: (qt.boundingBox.X1 + qt.boundingBox.X2) / 2,
		Y1: (qt.boundingBox.Y1 + qt.boundingBox.Y2) / 2, Y2: qt.boundingBox.Y2,
	}, points: map[Point]struct{}{}}

	qt.northWest = &QuadTree{boundingBox: &BoundingBox{
		X1: (qt.boundingBox.X1 + qt.boundingBox.X2) / 2, X2: qt.boundingBox.X2,
		Y1: (qt.boundingBox.Y1 + qt.boundingBox.Y2) / 2, Y2: qt.boundingBox.Y2,
	}, points: map[Point]struct{}{}}

	for p := range qt.points {
		qt.northEast.InsertPoint(&p)
		qt.northWest.InsertPoint(&p)
		qt.southEast.InsertPoint(&p)
		qt.southWest.InsertPoint(&p)
	}

	qt.points = nil
}

// TODO: fix this shitty naming
func (qt *QuadTree) exportPoints(receiver map[Point]struct{}) {
	if qt.points != nil {
		for p := range qt.points {
			receiver[p] = struct{}{}
		}
	}

	if qt.northEast != nil {
		qt.northEast.exportPoints(receiver)
		qt.northWest.exportPoints(receiver)
		qt.southEast.exportPoints(receiver)
		qt.southWest.exportPoints(receiver)
	}
}

func (qt *QuadTree) assemble() {
	qt.points = map[Point]struct{}{}

	qt.northEast.exportPoints(qt.points)
	qt.northWest.exportPoints(qt.points)
	qt.southEast.exportPoints(qt.points)
	qt.southWest.exportPoints(qt.points)

	qt.northEast = nil
	qt.northWest = nil
	qt.southEast = nil
	qt.southWest = nil
}

func (qt *QuadTree) InsertPoint(point *Point) {
	if !qt.boundingBox.Contains(point) {
		return
	}

	qt.size++

	if qt.points != nil && len(qt.points) >= nodeCapacity {
		qt.subdivide()
	}

	if qt.points == nil {
		qt.northEast.InsertPoint(point)
		qt.northWest.InsertPoint(point)
		qt.southEast.InsertPoint(point)
		qt.southWest.InsertPoint(point)
	} else {
		qt.points[*point] = struct{}{}
	}
}

func (qt *QuadTree) DeletePoint(point *Point) {
	if !qt.boundingBox.Contains(point) {
		return
	}

	qt.size--

	if qt.points == nil {
		qt.northEast.DeletePoint(point)
		qt.northWest.DeletePoint(point)
		qt.southEast.DeletePoint(point)
		qt.southWest.DeletePoint(point)
	} else {
		delete(qt.points, *point)
	}

	if qt.points == nil && len(qt.points) < nodeCapacity/2 {
		qt.assemble()
	}
}

func (qt *QuadTree) QueryClosestInRange(bb *BoundingBox, center *Point) *Point {
	if !qt.boundingBox.Intersects(bb) || qt.size == 0 {
		return nil
	}

	var result *Point
	if qt.points == nil {
		result = closest(center, result, qt.northEast.QueryClosestInRange(bb, center))
		result = closest(center, result, qt.northWest.QueryClosestInRange(bb, center))
		result = closest(center, result, qt.southEast.QueryClosestInRange(bb, center))
		result = closest(center, result, qt.southWest.QueryClosestInRange(bb, center))
	} else {
		for p := range qt.points {
			if bb.Contains(&p) {
				pcopy := p
				result = closest(center, result, &pcopy)
			}
		}
	}

	return result
}
