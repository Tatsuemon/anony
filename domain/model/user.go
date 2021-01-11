package model

import (
	"fmt"

	"github.com/google/uuid"
)

// User is CLI User.
type User struct {
	ID            string `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	Email         string `json:"email" db:"email"`
	EncryptedPass string `db:"password"`
}

// NewUser create a new user.
func NewUser(name string, email string, password string) (*User, error) {
	// ここでのバリデーションは入力された値に関するもの
	// DBにアクセスして他との比較調査するものはserviceにおく
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if password == "" {
		return nil, fmt.Errorf("password is required")
	}
	return &User{
		ID:            uuid.New().String(),
		Name:          name,
		Email:         email,
		EncryptedPass: password,
	}, nil
}

func EncryptedPassword(password string) (string, error) {
	// TODO(Tatsuemon): 暗号化したパスワードを返す
	return "encrypted password", nil
}

func ConfirmPassword(pass string, confirmPass string) (bool, error) {
	// TODO(Tatsuemon): パスワードの確認とバリデーション
	// これは, NewUserの前に行う
	// ConfirmPassword -> EncryptedPassword -> NewUser

	return true, nil
}
