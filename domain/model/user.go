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
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}, nil
}
