package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hiepphatle1104/nextmesh/internal/model"
	"github.com/hiepphatle1104/nextmesh/internal/utils"
)

func (s *Service) CreateAccount(ctx context.Context, data *model.RegisterDTO) error {

	if ok := utils.IsValidField(data.Name); !ok {
		return errors.New(model.ErrorInvalidName)
	}

	if ok := utils.IsValidField(data.Email); !ok {
		return errors.New(model.ErrorInvalidEmail)
	}

	if ok := utils.IsValidField(data.Password); !ok {
		return errors.New(model.ErrorInvalidPassword)
	}

	// Create account if valid
	account := &model.Account{
		ID:       uuid.NewString(),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	// Hash password before saving
	hashed, err := utils.HashPassword(account.Password)
	if err != nil {
		return err
	}

	account.Password = hashed
	return s.store.SaveAccount(ctx, account)
}
