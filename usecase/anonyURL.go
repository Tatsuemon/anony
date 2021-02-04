package usecase

import (
	"context"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
)

// AnonyURLUseCase is a usecase
type AnonyURLUseCase interface {
	CreateAnonyURL(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error)
}

type anonyURLUseCase struct {
	repo        repository.AnonyURLRepository
	transaction datastore.Transaction
	service     service.AnonyURLService
}

// NewAnonyURLUseCase creates conversionURLUseCase
func NewAnonyURLUseCase(r repository.AnonyURLRepository, t datastore.Transaction, s service.AnonyURLService) AnonyURLUseCase {
	return &anonyURLUseCase{r, t, s}
}

func (u *anonyURLUseCase) CreateAnonyURL(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error) {
	if err := u.service.IsExistedOriginalInUser(an.Original, userID); err != nil {
		return nil, err
	}
	if err := u.service.IsDuplicatedID(an.ID); err != nil {
		return nil, err
	}
	if err := an.ValidateAnonyURL(); err != nil {
		return nil, err
	}

	v, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		return u.repo.Save(ctx, an, userID)
	})
	if err != nil {
		return nil, err
	}
	return v.(*model.AnonyURL), nil
}
