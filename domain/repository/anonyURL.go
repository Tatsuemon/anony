package repository

import (
	"context"

	"github.com/Tatsuemon/anony/domain/model"
)

// AnonyURLRepository is a interface
type AnonyURLRepository interface {
	// ここには参照系と更新系のみのせる, 複数のmodelをまたがる場合は, usecase/queryserviceにかく
	FindByID(id string) (*model.AnonyURL, error)
	FindByUserID(userID string) ([]*model.AnonyURL, error)
	FindByUserIDWithStatus(userID string, status int64) ([]*model.AnonyURL, error)
	FindByOriginalInUser(original string, userID string) (*model.AnonyURL, error)
	Save(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error)
	UpdateStatus(ctx context.Context, an *model.AnonyURL) (*model.AnonyURL, error)
}
