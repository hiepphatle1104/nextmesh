package service

import (
	"context"
	"errors"

	"github.com/hiepphatle1104/nextmesh/internal/model"
)

func (s *Service) DeleteSession(ctx context.Context, sessionID string) error {
	session, err := s.store.FindSession(ctx, map[string]string{"id": sessionID})
	if err != nil {
		return err
	}

	if session == nil {
		return errors.New(model.ErrorSessionNotFound)
	}

	return s.store.DeleteSession(ctx, map[string]string{"id": sessionID}, session)
}
