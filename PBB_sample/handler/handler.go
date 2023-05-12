package handler

import (
	"errors"
)

type Handler struct {
	timeGen string
}

type Car struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

func (h *Handler) Handle(car Car) (Car, error) {
	if car.ID != "0" {
		return car, nil
	} else {
		return car, errors.New("no ID found")
	}
}
