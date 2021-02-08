package middleware

import (
	"context"
	"fmt"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/service"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type AuthFunc func(ctx context.Context, fullMethodName string) (context.Context, error)

func JWTAuth(s service.UserService) AuthFunc {
	return func(ctx context.Context, fullMethodName string) (context.Context, error) {
		if isSkipMethod(fullMethodName) {
			return ctx, nil
		}
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		auth, err := model.ParseJWT(token)
		if err != nil {
			return nil, err
		}
		userID := auth.UserID
		// ここで認可の処理を行う, userIDからUserIDを取得して, userが存在するかをみる
		flag, err := s.ExistsID(userID)
		if err != nil {
			return nil, err
		}
		if !flag {
			return nil, fmt.Errorf("token is invalid")
		}
		ctx = model.SetUserIDInContext(ctx, userID)
		return ctx, nil
	}
}

func isSkipMethod(method string) bool {
	// JWT認可がいらないMethodをここにかく
	skipMethodsName := [2]string{
		"/anony.UserService/CreateUser",
		"/anony.UserService/LogInUser",
	}

	for _, v := range skipMethodsName {
		if v == method {
			return true
		}
	}
	return false
}
