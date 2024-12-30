package repository

import (
	"context"

	"github.com/eiei114/go-backend-template/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, id, authToken, name string) error
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUserByUserId(ctx context.Context, id string) (domain.User, error)
	GetUserByAuthToken(ctx context.Context, authToken string) (*domain.User, error)
}
