package storage

import (
	"context"

	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func (s *Storage) SaveSession(ctx context.Context, session *model.Session) error {
	return s.db.WithContext(ctx).Create(session).Error
}
