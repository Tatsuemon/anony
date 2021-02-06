package model

import (
	"context"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// Auth is a struct for Authentification
type Auth struct {
	UserID   string
	UserName string
	Iss      string
	Iat      int64
	Exp      int64
}

// TODO(Tatsuemon): tokenの種類, JWTでやってみる

// NewJWT is
func NewJWT(userID, userName string, now time.Time) (string, error) {
	// JWT Tokenの作成場所
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"name": userName,
		"iss":  os.Getenv("JWT_ISS"),
		"iat":  now.Unix(),
		"exp":  now.Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
}

// ParseJWT is　はjwtから認証情報を取り出す
func ParseJWT(signed string) (*Auth, error) {
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		// ここでnoneを禁止している
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.Wrapf(err, "%s is expired", signed)
			}
		} else {
			return nil, errors.Wrapf(err, "%s is invalid", signed)
		}
	}

	if token == nil {
		return nil, errors.Errorf("not found token in %s:", signed)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.Errorf("not found claims in %s", signed)
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.Errorf("not found sub in %s", signed)
	}

	userName, ok := claims["name"].(string)
	if !ok {
		return nil, errors.Errorf("not found name in %s", signed)
	}

	iss, ok := claims["iss"].(string)
	if !ok {
		return nil, errors.Errorf("not found iss in %s", signed)
	}

	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, errors.Errorf("not found iat in %s", signed)
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.Errorf("not found exp in %s", signed)
	}

	return &Auth{
		UserID:   userID,
		UserName: userName,
		Iss:      iss,
		Iat:      int64(iat),
		Exp:      int64(exp),
	}, nil

}

type contextKey string

const userIDContextKey contextKey = "user_id"

// SetUserIDInContext set user_id in context
func SetUserIDInContext(parents context.Context, t string) context.Context {
	return context.WithValue(parents, userIDContextKey, t)
}

// GetUserIDInContext get user_id in context
func GetUserIDInContext(ctx context.Context) (string, error) {
	v := ctx.Value(userIDContextKey)
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("token not found")
	}
	return token, nil
}
