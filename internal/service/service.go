package service

import (
	"github.com/hiepphatle1104/nextmesh/internal/model"
)

type Service struct {
	store model.Storage
}

func NewService(store model.Storage) *Service {
	return &Service{store}
}
