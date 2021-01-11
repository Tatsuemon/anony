package usecase

import (
	"context"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
)

// UserUseCase is a usecase of user.
type UserUseCase interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userUseCase struct {
	repository.UserRepository
	transaction datastore.Transaction
}

// NewUserUseCase creates userUseCase.
func NewUserUseCase(r repository.UserRepository, t datastore.Transaction) UserUseCase {
	return &userUseCase{r, t}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	v, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		return u.UserRepository.Store(ctx, user)
	})
	if err != nil {
		return nil, err
	}

	return v.(*model.User), err
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// user.IDに目的のuserのIDが入っていることを期待する
	v, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		if _, err := u.UserRepository.FindByID(user.ID); err != nil {
			return nil, err
		}
		return u.UserRepository.Update(ctx, user)
	})
	if err != nil {
		return nil, err
	}

	return v.(*model.User), err
}

func (u *userUseCase) DeleteUser(ctx context.Context, id string) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		user, err := u.UserRepository.FindByID(id)
		if err != nil {
			return nil, err
		}
		return nil, u.UserRepository.Delete(ctx, user)
	})

	return err
}
