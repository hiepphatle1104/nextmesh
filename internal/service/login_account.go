package service

import (
	"context"
	"errors"

	"github.com/hiepphatle1104/nextmesh/internal/model"
	"github.com/hiepphatle1104/nextmesh/internal/utils"
)

func (s *Service) LoginAccount(ctx context.Context, data *model.LoginDTO) (*model.Account, error) {
	if ok := utils.IsValidField(data.Email); !ok {
		return nil, errors.New(model.ErrorInvalidEmail)
	}

	if ok := utils.IsValidField(data.Password); !ok {
		return nil, errors.New(model.ErrorInvalidPassword)
	}

	// Find account by email
	account, err := s.store.FindAccount(ctx, map[string]string{"email": data.Email})
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.New(model.ErrorAccountNotFound)
	}

	// Check password
	if err := utils.ComparePassword(account.Password, data.Password); err != nil {
		return nil, errors.New(model.ErrorIncorrectPassword)
	}

	return account, nil
}
