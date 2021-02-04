package model

import "fmt"

// ConversionURL is a conversion of url
type ConversionURL struct {
	ID       string
	Original string
	Short    string
	Status   int64 // 1: 待機中, 2: 有効, 3: 無効
}

// NewConversionURL create a new ConversionURL
func NewConversionURL(id string, original string, short string, status int64) (*ConversionURL, error) {
	return &ConversionURL{
		ID:       id,
		Original: original,
		Short:    short,
		Status:   status,
	}, nil
}

// ValidateConversionURL validates ConversionURL params
func ValidateConversionURL(id string, original string, short string, status int64) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	if original == "" {
		return fmt.Errorf("original is required")
	}
	if short == "" {
		return fmt.Errorf("short is required")
	}
	if status == 0 {
		return fmt.Errorf("status is required")
	}

	if (status >= 1) && (status <= 3) {
		return fmt.Errorf("status is out of range")
	}
	return nil
}
