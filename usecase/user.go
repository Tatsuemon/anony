package usecase

import (
	"context"
	"fmt"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"github.com/pkg/errors"
)

// UserUseCase is a usecase of user.
type UserUseCase interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	CheckDuplicatedUser(ctx context.Context, user *model.User) (bool, error)
	VerifyByNameOrEmailPass(ctx context.Context, nameOrEmail, password string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type userUseCase struct {
	repo        repository.UserRepository
	transaction datastore.Transaction
	service     service.UserService
}

// NewUserUseCase creates userUseCase.
func NewUserUseCase(r repository.UserRepository, t datastore.Transaction, s service.UserService) UserUseCase {
	return &userUseCase{r, t, s}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	exists, err := u.service.ExistsID(user.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("ID already existed")
	}
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	_, err = u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		return nil, u.repo.Save(ctx, user)
	})
	if err != nil {
		return nil, err
	}
	return u.repo.FindByID(user.ID)
}

// true: 重複するものは存在しない
func (u *userUseCase) CheckDuplicatedUser(ctx context.Context, user *model.User) (bool, error) {
	exist, err := u.service.ExistsDuplicatedUser(user.Name, user.Email)
	return !exist, err
}

func (u *userUseCase) VerifyByNameOrEmailPass(ctx context.Context, nameOrEmail, password string) (*model.User, error) {
	user, err := u.repo.FindByNameOrEmail(nameOrEmail)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.VerifyByNameOrEmailPass")
	}
	if ok := user.MatchPassword(password); !ok {
		return nil, fmt.Errorf("Wrong name or email, password")
	}
	return user, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// user.IDに目的のuserのIDが入っていることを期待する

	// TODO(Tatsuemon): nameとemailの重複を避ける
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		if _, err := u.repo.FindByID(user.ID); err != nil {
			return nil, err
		}
		return nil, u.repo.Update(ctx, user)
	})
	if err != nil {
		return nil, err
	}
	return u.repo.FindByID(user.ID)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id string) error {
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		user, err := u.repo.FindByID(id)
		if err != nil {
			return nil, err
		}
		return nil, u.repo.Delete(ctx, user)
	})
	return err
}
