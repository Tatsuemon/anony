package repository

import (
	"context"

	"github.com/Tatsuemon/anony/domain/model"
)

// UserRepository is a interface of UserRepository.
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Store(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, user *model.User) error
}
