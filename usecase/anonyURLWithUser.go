package usecase

import (
	"context"

	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"github.com/Tatsuemon/anony/usecase/queryservice"

	"github.com/Tatsuemon/anony/usecase/dto"
)

// AnonyURLWithUserUseCase is a usecase
type AnonyURLWithUserUseCase interface {
	CountByUser(ctx context.Context, userID string) (*dto.AnonyURLCountByUser, error)
}

type anonyURLWithUserUseCase struct {
	accessor    queryservice.UserAnonyURLAccessor
	transaction datastore.Transaction
}

// NewAnonyURLWithUserUseCase creates anonyURLWithUserUsecase
func NewAnonyURLWithUserUseCase(a queryservice.UserAnonyURLAccessor, t datastore.Transaction) AnonyURLWithUserUseCase {
	return &anonyURLWithUserUseCase{a, t}
}

func (a anonyURLWithUserUseCase) CountByUser(ctx context.Context, userID string) (*dto.AnonyURLCountByUser, error) {
	return a.accessor.CountAnonyURLByUser(userID)
}
