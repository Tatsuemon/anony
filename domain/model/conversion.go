package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// Conversion is a set of urls
type Conversion struct {
	ShortURL    string `json:"short_url" db:"short_url"`
	OriginalURL string `json:"original_url" db:"original_url"`
	Status      int32  `json:"status" db:"status"` // 0: 非表示, 1: 表示
}

// NewConversion creates a new Conversion
func NewConversion(shortURL, originalURL string) (*Conversion, error) {

	validURL := govalidator.IsURL(shortURL)
	if !validURL {
		return nil, errors.New("shortURL is not valid url.")
	}

	validURL = govalidator.IsURL(originalURL)
	if !validURL {
		return nil, errors.New("originalURL is not valid url.")
	}

	return &Conversion{
		ShortURL:    shortURL,
		OriginalURL: originalURL,
		Status:      1,
	}, nil
}
