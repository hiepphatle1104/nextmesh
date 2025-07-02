package model

import "context"

type Storage interface {
	SaveAccount(context.Context, *Account) error
	FindAccount(context.Context, map[string]string) (*Account, error)
	SaveSession(context.Context, *Session) error
	FindSession(context.Context, map[string]string) (*Session, error)
	DeleteSession(context.Context, map[string]string, *Session) error
}

type Service interface {
	LoginAccount(context.Context, *LoginDTO) (*Account, error)
	GetSession(context.Context, string) (*Session, error)
	CreateSession(context.Context, string) (*Session, error)
	CreateAccount(context.Context, *RegisterDTO) error
	DeleteSession(context.Context, string) error
}
