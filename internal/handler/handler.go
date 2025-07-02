package handler

import (
	"github.com/hiepphatle1104/nextmesh/internal/model"
)

type Controller struct {
	service model.Service
}

func NewController(service model.Service) *Controller {
	return &Controller{service}
}
