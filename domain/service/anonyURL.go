package service

import (
	"errors"

	"github.com/Tatsuemon/anony/domain/repository"
)

// AnonyURLService is a service.
type AnonyURLService interface {
	IsDuplicatedID(id string) error
	IsExistedOriginalInUser(original string, userID string) error
}

type anonyURLService struct {
	repo repository.AnonyURLRepository
}

// NewAnonyURLService create a new service.
func NewAnonyURLService(r repository.AnonyURLRepository) AnonyURLService {
	return &anonyURLService{r}
}

func (c *anonyURLService) IsDuplicatedID(id string) error {
	cURL, err := c.repo.FindByID(id) // TODO(Tatsuemon): sql.NoRowみたいなエラーが出ない方法で行う
	if err != nil {
		return err
	}
	if len(cURL) != 0 {
		err = errors.New("id is duplicated")
		return err
	}
	return nil
}

func (c *anonyURLService) IsExistedOriginalInUser(original string, userID string) error {
	cURL, err := c.repo.FindByOriginalInUser(original, userID)
	if err != nil {
		return err
	}
	if len(cURL) != 0 {
		err = errors.New("You have already created this original url")
		return err
	}
	return nil
}
