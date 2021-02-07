package handler

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/google/uuid"

	"github.com/Tatsuemon/anony/rpc"
	"github.com/Tatsuemon/anony/usecase"
)

// UserHandler implements rpc.UserServiceServer interface
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
	// CLIで制御できる値
	/*
		- name, emailが重複している
	*/
	name := in.GetUser().GetName()
	email := in.GetUser().GetEmail()
	password := in.GetPassword()
	confirmPassword := in.GetConfirmPassword()

	if _, err := model.ConfirmPassword(password, confirmPassword); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "password is invalid \n: %s", err)
	}

	encryptedPass, err := model.EncryptPassword(password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to encrypted password \n: %s", err)
	}

	id := uuid.New().String()
	user := model.NewUser(id, name, email, encryptedPass)

	// TODO(Tatsuemon): 重複処理について
	ok, err := u.UserUseCase.CheckDuplicatedUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to check duplication user \n: %s", err)
	}
	if !ok {
		return nil, status.Errorf(codes.AlreadyExists, "name or email is already exists")
	}

	user, err = u.UserUseCase.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create user \n: %s", err)
	}

	// JWT Tokenの作成
	token, err := model.NewJWT(user.ID, user.Name, time.Now())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create JWT \n: %s", err)
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

// LogInUser log in to anony using user infomation.
func (u *UserHandler) LogInUser(ctx context.Context, in *rpc.LogInUserRequest) (*rpc.LogInUserResponse, error) {
	nameOrEmail := in.GetNameOrEmail()
	password := in.GetPassword()

	user, err := u.UserUseCase.VerifyByNameOrEmailPass(ctx, nameOrEmail, password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to log in to anony \n:%s", err)
	}

	// JWT Tokenの作成
	token, err := model.NewJWT(user.ID, user.Name, time.Now())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create JWT \n: %s", err)
	}

	res := &rpc.LogInUserResponse{
		User: &rpc.UserBase{
			Name:  user.Name,
			Email: user.Email,
		},
		Token: token,
	}
	return res, nil
}
