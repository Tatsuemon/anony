package service

import (
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
}

type userService struct {
	repository.UserRepository
}

// NewUserService create a new service of user.
func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (u *userService) ExistsID(id string) (bool, error) {
	_, err := u.UserRepository.FindByID(id)
	if err != nil {
		return false, errors.Wrap(err, "failed to userService.ExistsID")
	}
	return true, nil
}

func (u *userService) ExistsName(name string) (bool, error) {
	_, err := u.UserRepository.FindByName(name)
	if err != nil {
		return false, errors.Wrap(err, "failed to userService.ExistsName")
	}
	return true, nil
}
func (u *userService) ExistsEmail(email string) (bool, error) {
	_, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return false, errors.Wrap(err, "failed to userService.ExistsEmail")
	}
	return true, nil
}
