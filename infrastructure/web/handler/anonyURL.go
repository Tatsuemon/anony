package handler

import (
	"context"
	"log"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/rpc"
	"github.com/Tatsuemon/anony/usecase"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AnonyURLHandler implements rpc.AnonyURLService interface
type AnonyURLHandler struct {
	usecase         usecase.AnonyURLUseCase
	usecaseWithUser usecase.AnonyURLWithUserUseCase
}

// NewAnonyURLHandler creates a new UserHandler
func NewAnonyURLHandler(u usecase.AnonyURLUseCase, uu usecase.AnonyURLWithUserUseCase) *AnonyURLHandler {
	return &AnonyURLHandler{u, uu}
}

// CreateAnonyURL creates anonyURL
func (a *AnonyURLHandler) CreateAnonyURL(ctx context.Context, in *rpc.CreateAnonyURLRequest) (*rpc.CreateAnonyURLResponse, error) {
	userID, err := model.GetUserIDInContext(ctx)
	if err != nil {
		return nil, err
	}

	ori := in.GetOriginalUrl()
	isActive := in.GetIsActive()
	var status int64
	if isActive {
		status = 1
	} else {
		status = 2
	}
	su, err := a.usecase.CreateAnonyURL(ctx, userID)
	if err != nil {
		return nil, err
	}
	an := model.NewAnonyURL(uuid.New().String(), ori, su, status)
	_, err = a.usecase.SaveAnonyURL(ctx, an, userID)
	if err != nil {
		return nil, err
	}

	res := &rpc.CreateAnonyURLResponse{
		AnonyUrls: &rpc.AnonyURL{
			OriginalUrl: an.Original,
			ShortUrl:    an.Short,
			IsActive:    an.Status == 1,
		},
	}
	return res, nil
}

// ListAnonyURLs lists user's Anony URLs
func (a *AnonyURLHandler) ListAnonyURLs(ctx context.Context, in *rpc.ListAnonyURLsRequest) (*rpc.ListAnonyURLsResponse, error) {
	userID, err := model.GetUserIDInContext(ctx)
	if err != nil {
		return nil, err
	}
	inActive := in.GetInActive()
	all := in.GetAll()
	var status int64
	if all {
		status = 0
	} else if inActive {
		status = 2
	} else {
		status = 1
	}

	ans, err := a.usecase.ListAnonyURLs(ctx, userID, status)
	if err != nil {
		return nil, err
	}

	res := &rpc.ListAnonyURLsResponse{}
	res.AnonyUrls = make([]*rpc.AnonyURL, len(ans))
	for i, v := range ans {
		res.AnonyUrls[i] = &rpc.AnonyURL{
			OriginalUrl: v.Original,
			ShortUrl:    v.Short,
			IsActive:    v.Status == 1,
		}
	}
	return res, nil
}

// CountAnonyURLs count user's anony urls
func (a *AnonyURLHandler) CountAnonyURLs(ctx context.Context, in *emptypb.Empty) (*rpc.CountAnonyURLsResponse, error) {
	userID, err := model.GetUserIDInContext(ctx)
	if err != nil {
		return nil, err
	}
	ans, err := a.usecaseWithUser.CountByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := &rpc.CountAnonyURLsResponse{
		Name:        ans.Name,
		Email:       ans.Email,
		CountAll:    ans.CntURLs,
		CountActive: ans.CntActiveURLs,
	}
	return res, nil
}

// UpdateAnonyURLStatus change status of AnonyURL
func (a *AnonyURLHandler) UpdateAnonyURLStatus(ctx context.Context, in *rpc.UpdateAnonyURLStatusRequest) (*rpc.UpdateAnonyURLStatusResponse, error) {
	userID, err := model.GetUserIDInContext(ctx)
	if err != nil {
		return nil, err
	}
	ori := in.GetOriginalUrl()
	isActive := in.GetIsActive()
	var status int64
	if isActive {
		status = 1
	} else {
		status = 2
	}
	ans, err := a.usecase.UpdateAnonyURLStatus(ctx, ori, userID, status)
	if err != nil {
		return nil, err
	}
	log.Println(ans)
	res := &rpc.UpdateAnonyURLStatusResponse{
		AnonyUrl: &rpc.AnonyURL{
			OriginalUrl: ans.Original,
			ShortUrl:    ans.Short,
			IsActive:    ans.Status == 1,
		},
	}

	return res, nil
}
