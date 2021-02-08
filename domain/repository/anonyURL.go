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
	FindByAnonyURL(anonyURL string) (*model.AnonyURL, error)
	GetIDByOriginalUser(original, userID string) (string, error)
	Save(ctx context.Context, an *model.AnonyURL, userID string) error
	UpdateStatus(ctx context.Context, id string, status int64) error
}
