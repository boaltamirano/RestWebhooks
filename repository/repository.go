package repository

import (
	"context"

	"github.com/RestWebkooks/models"
)

type Repository interface {
	// User
	InsetUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// Post
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)

	// Close conection
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

// **********************************USERS********************************************** //
func InsetUser(ctx context.Context, user *models.User) error {
	return implementation.InsetUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

// **********************************POSTS********************************************** //
func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
