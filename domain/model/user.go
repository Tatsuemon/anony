package model

import (
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User is CLI User.
type User struct {
	ID            string `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	Email         string `json:"email" db:"email"`
	EncryptedPass string `db:"password"`
}

// NewUser create a new user.
func NewUser(id string, name string, email string, password string) (*User, error) {
	return &User{
		ID:            id,
		Name:          name,
		Email:         email,
		EncryptedPass: password,
	}, nil
}

// ValidateUser validates User params
func (u *User) ValidateUser() error {
	if u.ID == "" {
		return fmt.Errorf("id is required")
	}
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.EncryptedPass == "" {
		return fmt.Errorf("password is required")
	}
	if utf8.RuneCountInString(u.EncryptedPass) < 6 && utf8.RuneCountInString(u.EncryptedPass) > 0 {
		return fmt.Errorf("password must be at least 6 characters")
	}
	return nil
}

// MatchPassword returns whether it matches encrypted password
func (u *User) MatchPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPass), []byte(password)) == nil
}

// EncryptPassword encrypt password
func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "failed to EncryptPassword")
	}
	return string(hash), nil
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
