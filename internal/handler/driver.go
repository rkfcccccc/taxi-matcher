package handler

import (
	"net/http"

	"github.com/rkfcccccc/taxi_matcher/internal/service"
)

type Driver struct {
	Id int `json:"id"`
	X  int `json:"x"`
	Y  int `json:"y"`
}

func (h *Handler) Driver(w http.ResponseWriter, req *http.Request) {
	var input Driver
	if err := h.bindJSON(w, req, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case "GET":
		pedestrian := &service.Pedestrian{X: input.X, Y: input.Y}
		driver := h.service.SearchClosest(pedestrian)

		if driver == nil {
			h.writeJSON(w, http.StatusNotFound, map[string]interface{}{
				"error": "DRIVER_NOT_FOUND",
			})
		} else {
			h.writeJSON(w, http.StatusNotFound, Driver{driver.Id, driver.X, driver.Y})
		}
	case "POST":
		driver := &service.Driver{Id: input.Id, X: input.X, Y: input.Y}
		h.service.Add(driver)
	case "DELETE":
		driver := &service.Driver{Id: input.Id, X: input.X, Y: input.Y}
		h.service.Delete(driver)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
