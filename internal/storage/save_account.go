package storage

import (
	"context"

	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func (s *Storage) SaveAccount(ctx context.Context, data *model.Account) error {
	return s.db.WithContext(ctx).Create(&data).Error
}
