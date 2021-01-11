package service

import (
	"database/sql"

	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/pkg/errors"
)

// これらをusecaseで使用する Storeのとき
// TODO(Tatsuemon): errの加工

// UserService is a service of User.
type UserService interface {
	ExistsID(id string) (bool, error)
	ExistsName(name string) (bool, error)
	ExistsEmail(email string) (bool, error)
	ExistsParams(id, name, email string) (bool, error)
}

type userService struct {
	repository.UserRepository
}

// NewUserService create a new service of user.
func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (u *userService) ExistsID(id string) (bool, error) {
	user, err := u.UserRepository.FindByID(id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "failed to userService.ExistsID")
	}
	return user != nil, nil
}

func (u *userService) ExistsName(name string) (bool, error) {
	user, err := u.UserRepository.FindByName(name)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "failed to userService.ExistsName")
	}
	return user != nil, nil
}

func (u *userService) ExistsEmail(email string) (bool, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "failed to userService.ExistsEmail")
	}
	return user != nil, nil
}

func (u *userService) ExistsParams(id, name, email string) (bool, error) {
	exists, err := u.ExistsID(id)
	if err != nil {
		return false, err
	}
	if exists {
		return true, errors.New("this user ID is already existed")
	}

	exists, err = u.ExistsName(name)
	if err != nil {
		return false, err
	}
	if exists {
		return true, errors.New("this user name is already existed")
	}

	exists, err = u.ExistsEmail(email)
	if err != nil {
		return false, err
	}
	if exists {
		return true, errors.New("this user email is already existed")
	}
	return false, nil
}
