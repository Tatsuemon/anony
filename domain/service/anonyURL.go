package service

import (
	"github.com/Tatsuemon/anony/domain/repository"
)

// AnonyURLService is a service.
type AnonyURLService interface {
	ExistID(id string) (bool, error)
	ExistOriginalInUser(original, userID string) (bool, error)
}

type anonyURLService struct {
	repo repository.AnonyURLRepository
}

// NewAnonyURLService create a new service.
func NewAnonyURLService(r repository.AnonyURLRepository) AnonyURLService {
	return &anonyURLService{r}
}

func (a *anonyURLService) ExistID(id string) (bool, error) {
	an, err := a.repo.FindByID(id)
	if err != nil {
		return false, err
	}
	return an != nil, nil
}

func (a *anonyURLService) ExistOriginalInUser(original, userID string) (bool, error) {
	an, err := a.repo.FindByOriginalInUser(original, userID)
	if err != nil {
		return false, err
	}
	return an != nil, nil
}
