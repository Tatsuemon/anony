package middleware

import (
	"context"
	"fmt"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type AuthFunc func(ctx context.Context, fullMethodName string) (context.Context, error)

func JWTAuth() AuthFunc {
	return func(ctx context.Context, fullMethodName string) (context.Context, error) {
		if isSkipMethod(fullMethodName) {
			return ctx, nil
		}
		// TODO(Tatsuemon): ここで認可の処理を行う

		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		fmt.Print(token)
		if err != nil {
			return nil, err
		}
		fmt.Print(token)

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
