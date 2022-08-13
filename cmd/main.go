package main

import (
	"log"
	"net/http"

	"github.com/rkfcccccc/taxi_matcher/internal/handler"
	"github.com/rkfcccccc/taxi_matcher/internal/service"
)

const dimension int = 1e5

func main() {
	service := service.NewService(dimension)
	handler := handler.NewHandler(service)

	http.HandleFunc("/driver", handler.Driver)

	log.Println("Server started")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
