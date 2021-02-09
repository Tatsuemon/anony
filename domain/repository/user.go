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
	FindByNameOrEmail(nameOrEmail string) (*model.User, error)
	FindDuplicatedUsers(name, email string) ([]*model.User, error)
	Save(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}
