package service

import "testing"

func TestService(t *testing.T) {
	service := NewService(10)

	d1 := &Driver{Id: 0, X: 0, Y: 0}
	d2 := &Driver{Id: 1, X: 9, Y: 9}

	service.Add(d1)
	service.Add(d2)

	result := service.SearchClosest(&Pedestrian{X: 2, Y: 2})
	if result == nil || *result != *d1 {
		t.Fatalf("found driver is invalid: %+v", result)
	}

	service.Delete(d1)

	result = service.SearchClosest(&Pedestrian{X: 3, Y: 3})
	if result == nil || *result != *d2 {
		t.Fatalf("found driver is invalid: %+v", result)
	}

	service.Delete(d2)

	result = service.SearchClosest(&Pedestrian{X: 4, Y: 4})
	if result != nil {
		t.Fatalf("found driver but shouldnt: %+v", result)
	}
}
