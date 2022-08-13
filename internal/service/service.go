package service

import (
	"sync"

	qtree "github.com/rkfcccccc/taxi_matcher/pkg/quadtree"
)

type Service struct {
	dimension int
	qtree     *qtree.QuadTree
	mutex     sync.RWMutex
}

func NewService(dimension int) *Service {
	return &Service{
		dimension: dimension,
		qtree:     qtree.NewQuadTree(dimension, dimension),
		mutex:     sync.RWMutex{},
	}
}

func (service *Service) Add(d *Driver) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	service.qtree.InsertPoint(d.toPoint())
}

func (service *Service) Delete(d *Driver) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	service.qtree.DeletePoint(d.toPoint())
}

func (service *Service) SearchClosest(pedestrian *Pedestrian) *Driver {
	service.mutex.RLock()
	defer service.mutex.RUnlock()

	point := pedestrian.toPoint()

	for i := 1; i <= service.dimension; i <<= 1 {
		boundings := &qtree.BoundingBox{
			X1: point.X - i, Y1: point.Y - i,
			X2: point.X + i, Y2: point.Y + i,
		}

		result := service.qtree.QueryClosestInRange(boundings, point)
		if result != nil {
			return driverFromPoint(result)
		}
	}

	return nil
}
