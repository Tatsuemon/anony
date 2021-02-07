package service

import (
	"fmt"

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
	an, err := c.repo.FindByID(id)
	if err != nil {
		return err
	}
	if an != nil {
		return fmt.Errorf("id is duplicated")
	}
	return nil
}

func (c *anonyURLService) IsExistedOriginalInUser(original string, userID string) error {
	an, err := c.repo.FindByOriginalInUser(original, userID)
	if err != nil {
		return err
	}
	if an != nil {
		return fmt.Errorf("You have already created this original url")
	}
	return nil
}
