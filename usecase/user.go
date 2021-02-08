package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
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
	repository.UserRepository
	transaction datastore.Transaction
	service.UserService
}

// NewUserUseCase creates userUseCase.
func NewUserUseCase(r repository.UserRepository, t datastore.Transaction, s service.UserService) UserUseCase {
	return &userUseCase{r, t, s}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	exists, err := u.UserService.ExistsID(user.ID)
	if err != nil || exists {
		return nil, err
	}
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	v, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		return u.UserRepository.Save(ctx, user)
	})
	if err != nil {
		return nil, err
	}
	return v.(*model.User), nil
}

// true: 重複するものは存在しない
func (u *userUseCase) CheckDuplicatedUser(ctx context.Context, user *model.User) (bool, error) {
	exist, err := u.ExistsDuplicatedUser(user.Name, user.Email)
	return !exist, err
}

func (u *userUseCase) VerifyByNameOrEmailPass(ctx context.Context, nameOrEmail, password string) (*model.User, error) {
	user, err := u.UserRepository.FindByNameOrEmail(nameOrEmail)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Wrong name or email, password")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to userRepository.FindByNameOrEmailPass")
	}
	if ok := user.MatchPassword(password); !ok {
		return nil, fmt.Errorf("Wrong name or email, password")
	}

	return user, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// user.IDに目的のuserのIDが入っていることを期待する

	// TODO(Tatsuemon): nameとemailの重複を避ける
	v, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		if _, err := u.UserRepository.FindByID(user.ID); err != nil {
			return nil, err
		}
		return u.UserRepository.Update(ctx, user)
	})
	if err != nil {
		return nil, err
	}

	return v.(*model.User), nil
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
