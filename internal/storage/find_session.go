package storage

import (
	"context"

	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func (s *Storage) FindSession(ctx context.Context, params map[string]string) (*model.Session, error) {
	var session model.Session
	if err := s.db.WithContext(ctx).Where(params).First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}
