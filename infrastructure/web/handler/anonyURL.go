package handler

import (
	"context"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/rpc"
	"github.com/Tatsuemon/anony/usecase"
	"github.com/google/uuid"
)

// AnonyURLHandler implements rpc.AnonyURLService interface
type AnonyURLHandler struct {
	usecase usecase.AnonyURLUseCase
}

// NewAnonyURLHandler creates a new UserHandler
func NewAnonyURLHandler(u usecase.AnonyURLUseCase) *AnonyURLHandler {
	return &AnonyURLHandler{u}
}

// CreateAnonyURL creates anonyURL
func (a *AnonyURLHandler) CreateAnonyURL(ctx context.Context, in *rpc.CreateAnonyURLRequest) (*rpc.CreateAnonyURLResponse, error) {
	ori := in.GetOriginalUrl()

	// TODO(Tatsuemon): JWTからuser_ID取得
	userID := "test"

	// TODO(Tatsuemon): shortURLのロジック
	su := "http://shorttttttt/"

	an, err := model.NewAnonyURL(uuid.New().String(), ori, su, 1)
	if err != nil {
		return nil, err
	}
	_, err = a.usecase.CreateAnonyURL(ctx, an, userID)
	if err != nil {
		return nil, err
	}

	res := &rpc.CreateAnonyURLResponse{
		OriginalUrl: an.Original,
		ShortUrl:    an.Short,
		Status:      an.Status,
	}

	return res, nil
}