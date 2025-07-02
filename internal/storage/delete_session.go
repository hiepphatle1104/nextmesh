package storage

import (
	"context"

	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func (s *Storage) DeleteSession(ctx context.Context, params map[string]string, session *model.Session) error {
	return s.db.WithContext(ctx).Where(params).Delete(session).Error
}
