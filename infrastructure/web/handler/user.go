package handler

import (
	"context"
	"errors"

	"github.com/Tatsuemon/anony/domain/model"

	"github.com/Tatsuemon/anony/rpc"
	"github.com/Tatsuemon/anony/usecase"
)

// UserHandler implements rpc.UserSErviceServer interface
type UserHandler struct {
	usecase.UserUseCase
}

// CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
// 	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
// 	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)

func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{u}
}

func (u *UserHandler) CreateUser(ctx context.Context, in *rpc.CreateUserRequest) (*rpc.CreateUserResponse, error) {
	name := in.GetUser().GetName()
	email := in.GetUser().GetEmail()
	password := in.GetPassword()
	confirmPassword := in.GetConfirmPassword()

	if _, err := model.ConfirmPassword(password, confirmPassword); err != nil {
		return nil, errors.New("password is invalid")
	}

	encryptedPass, err := model.EncryptedPassword(password)
	if err != nil {
		return nil, errors.New("failed to encrypted password")
	}

	user, err := model.NewUser(name, email, encryptedPass)
	if err != nil {
		return nil, errors.New("failed to NewUser")
	}

	user, err = u.UserUseCase.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// TODO(Tatsuemon): ログイン処理, 適当にtokenを作成と登録処理
	token := "token1"

	res := &rpc.CreateUserResponse{
		User: &rpc.UserBase{
			Name:  name,
			Email: email,
		},
		Token: token,
	}
	return res, nil
}
