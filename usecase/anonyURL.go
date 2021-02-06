package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"os"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
)

// AnonyURLUseCase is a usecase
type AnonyURLUseCase interface {
	CreateAnonyURL(ctx context.Context, userID string) (string, error)
	SaveAnonyURL(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error)
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

func (u *anonyURLUseCase) SaveAnonyURL(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error) {
	// TODO(Tatsuemon): すでにあって, statusが１でない場合は1に変更する
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

func (u *anonyURLUseCase) CreateAnonyURL(ctx context.Context, userID string) (string, error) {
	host := os.Getenv("SERVER_HOST")
	path := userID[25:33] + "/"

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}
	// letters からランダムに取り出して文字列を生成
	for _, v := range b {
		// index が letters の長さに収まるように調整
		path += string(letters[int(v)%len(letters)])
	}
	return host + "/" + path, nil
}
