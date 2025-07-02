package service

import (
	"context"
	"errors"

	"github.com/hiepphatle1104/nextmesh/internal/model"
	"github.com/hiepphatle1104/nextmesh/internal/utils"
)

func (s *Service) GetSession(ctx context.Context, sessionID string) (*model.Session, error) {
	session, err := s.store.FindSession(ctx, map[string]string{"id": sessionID})
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.New(model.ErrorSessionNotFound)
	}

	if ok := utils.CompareHash(session.ID, sessionID); !ok {
		return nil, errors.New(model.ErrorSessionInvalid)
	}

	return session, nil
}
