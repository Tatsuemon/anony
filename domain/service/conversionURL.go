package service

import (
	"errors"

	"github.com/Tatsuemon/anony/domain/repository"
)

// ConversionURLService is a service.
type ConversionURLService interface {
	IsDuplicatedID(id string) error
	IsExistedOriginalURLOfUser(original string, userID string) error
}

type conversionURLService struct {
	repo repository.ConversionURLRepository
}

func (c *conversionURLService) IsDuplicatedID(id string) error {
	cURL, err := c.repo.FindByID(id) // TODO(Tatsuemon): sql.NoRowみたいなエラーが出ない方法で行う
	if err != nil {
		return err
	}
	if cURL != nil {
		err = errors.New("id is duplicated")
		return err
	}
	return nil
}

func (c *conversionURLService) IsExistedOriginalURLOfUser(original string, userID string) error {
	cURL, err := c.repo.FindByOriginalURLOfUser(original, userID)
	if err != nil {
		return err
	}
	if cURL != nil {
		err = errors.New("You have already created this original url")
		return err
	}
	return nil
}
