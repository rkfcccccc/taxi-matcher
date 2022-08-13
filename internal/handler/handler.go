package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rkfcccccc/taxi_matcher/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) bindJSON(w http.ResponseWriter, req *http.Request, v any) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(bytes); err != nil {
		return err
	}

	return nil
}
