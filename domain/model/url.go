package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

// URL is URL info
// 値オブジェクト
type URL struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// NewURL creates a new URL
func NewURL(name string) (*URL, error) {
	validURL := govalidator.IsURL(name)
	if !validURL {
		return nil, errors.New("name is not valid url.")
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	uu := u.String()

	return &URL{
		ID:   uu,
		Name: name,
	}, nil
}
