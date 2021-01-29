package handler

import (
	"context"
	"errors"
	"time"

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

// NewUserHandler creates a new UserHandler
func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{u}
}

// CreateUser creates a new user
func (u *UserHandler) CreateUser(ctx context.Context, in *rpc.CreateUserRequest) (*rpc.CreateUserResponse, error) {
	name := in.GetUser().GetName()
	email := in.GetUser().GetEmail()
	password := in.GetPassword()
	confirmPassword := in.GetConfirmPassword()

	if _, err := model.ConfirmPassword(password, confirmPassword); err != nil {
		return nil, errors.New("password is invalid")
	}

	encryptedPass, err := model.EncryptPassword(password)
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

	// JWT Tokenの作成
	token, err := model.NewJWT(user.ID, user.Name, time.Now())
	if err != nil {
		return nil, errors.New("failed to Create JWT Token")
	}

	res := &rpc.CreateUserResponse{
		User: &rpc.UserBase{
			Name:  name,
			Email: email,
		},
		Token: token,
	}
	return res, nil
}
