package model

import (
	"fmt"
	"unicode/utf8"

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

// ConfirmPassword is check password
func ConfirmPassword(pass string, confirmPass string) (bool, error) {
	if pass == "" {
		return false, fmt.Errorf("password is required")
	}
	if confirmPass == "" {
		return false, fmt.Errorf("confirm_password is required")
	}

	if pass != confirmPass {
		return false, fmt.Errorf("password needs to equeal to confirm_password")
	}

	if utf8.RuneCountInString(pass) < 6 {
		return false, fmt.Errorf("password must be at least 6 characters")
	}

	return true, nil
}
