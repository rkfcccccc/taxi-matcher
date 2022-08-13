package service

import (
	qtree "github.com/rkfcccccc/taxi_matcher/pkg/quadtree"
)

type Driver struct {
	Id, X, Y int
}

type Pedestrian struct {
	X, Y int
}

func (driver *Driver) toPoint() *qtree.Point {
	return &qtree.Point{Key: driver.Id, X: driver.X, Y: driver.Y}
}

func (ped *Pedestrian) toPoint() *qtree.Point {
	return &qtree.Point{Key: -1, X: ped.X, Y: ped.Y}
}

func driverFromPoint(p *qtree.Point) *Driver {
	return &Driver{p.Key, p.X, p.Y}
}
