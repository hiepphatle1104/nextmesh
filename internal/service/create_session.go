package service

import (
	"context"
	"errors"
	"time"

	"github.com/hiepphatle1104/nextmesh/internal/model"
	"github.com/hiepphatle1104/nextmesh/internal/utils"
)

func (s *Service) CreateSession(ctx context.Context, accountID string) (*model.Session, error) {
	// Find session
	exist, err := s.store.FindSession(ctx, map[string]string{"account_id": accountID})
	if err != nil {
		return nil, err
	}

	if exist != nil {

		// TODO: Refresh session if needed
		return exist, nil
	}

	now := time.Now().UTC()
	hashedID, err := utils.GenerateHash(32)
	if err != nil {
		return nil, err
	}

	session := &model.Session{
		ID:        hashedID,
		ExpiresAt: now.Add(time.Minute * 5),
		AccountID: accountID,
		CreatedAt: now,
	}

	if err := s.store.SaveSession(ctx, session); err != nil {
		return nil, err
	}

	if ok := utils.IsValidField(session.ID); !ok {
		return nil, errors.New(model.ErrorCreateSession)
	}

	return session, nil
}
