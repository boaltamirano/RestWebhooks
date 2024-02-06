package repository

import (
	"context"

	"github.com/RestWebkooks/models"
)

type UserRepository interface {
	InsetUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsetUser(ctx context.Context, user *models.User) error {
	return implementation.InsetUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
