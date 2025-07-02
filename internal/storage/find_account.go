package storage

import (
	"context"

	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func (s *Storage) FindAccount(ctx context.Context, params map[string]string) (*model.Account, error) {
	var account model.Account
	if err := s.db.WithContext(ctx).Where(params).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
