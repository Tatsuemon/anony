package usecase

import (
	"context"
	"crypto/rand"
	"fmt"
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
	UpdateAnonyURLStatus(ctx context.Context, original, userID string, status int64) (*model.AnonyURL, error)
	ListAnonyURLs(ctx context.Context, userID string, q int64) ([]*model.AnonyURL, error)
	GetOriginalByAnonyURL(ctx context.Context, anonyURL string) (string, error)
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

func (u *anonyURLUseCase) CreateAnonyURL(ctx context.Context, userID string) (string, error) {
	host := os.Getenv("SERVER_HOST")
	path := userID[25:33] + "/"

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("unexpected error")
	}
	// letters からランダムに取り出して文字列を生成
	for _, v := range b {
		// index が letters の長さに収まるように調整
		path += string(letters[int(v)%len(letters)])
	}
	return host + "/" + path, nil
}

func (u *anonyURLUseCase) SaveAnonyURL(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error) {
	exist, err := u.service.ExistOriginalInUser(an.Original, userID)
	if err != nil {
		return nil, err
	}
	idExisted, err := u.service.ExistID(an.ID)
	if err != nil {
		return nil, err
	}
	if idExisted {
		return nil, fmt.Errorf("id is already existed")
	}

	if err := an.ValidateAnonyURL(); err != nil {
		return nil, err
	}
	_, err = u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		if exist {
			id, err := u.repo.GetIDByOriginalUser(an.Original, userID)
			if err != nil {
				return nil, err
			}
			an.ID = id
			return nil, u.repo.UpdateStatus(ctx, id, an.Status)
		}
		return nil, u.repo.Save(ctx, an, userID)
	})
	if err != nil {
		return nil, err
	}
	return u.repo.FindByID(an.ID)
}

func (u *anonyURLUseCase) UpdateAnonyURLStatus(ctx context.Context, original, userID string, status int64) (*model.AnonyURL, error) {
	if status < 1 || status > 2 {
		return nil, fmt.Errorf("status is out of range")
	}
	var id string
	_, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		aid, err := u.repo.GetIDByOriginalUser(original, userID)
		id = aid
		if err != nil {
			return nil, err
		}
		if id == "" {
			return nil, fmt.Errorf("this anonyURL is not existed")
		}
		return nil, u.repo.UpdateStatus(ctx, id, status)
	})
	if err != nil {
		return nil, err
	}
	return u.repo.FindByID(id)
}

func (u *anonyURLUseCase) ListAnonyURLs(ctx context.Context, userID string, q int64) ([]*model.AnonyURL, error) {
	// q=0 -> all
	// q=1 -> active
	// q=2 -> inactive
	if q == 0 {
		return u.repo.FindByUserID(userID)
	} else if q == 1 {
		return u.repo.FindByUserIDWithStatus(userID, 1)
	} else if q == 2 {
		return u.repo.FindByUserIDWithStatus(userID, 2)
	} else {
		return nil, fmt.Errorf("out of range")
	}
}

func (u *anonyURLUseCase) GetOriginalByAnonyURL(ctx context.Context, anonyURL string) (string, error) {
	an, err := u.repo.FindByAnonyURL(anonyURL)
	if err != nil {
		return "", err
	}
	if an == nil {
		return "", nil
	}
	if an.Status != 1 {
		return "", nil
	}
	return an.Original, nil
}
