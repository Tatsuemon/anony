package model

import "fmt"

// AnonyURL is a conversion of url
type AnonyURL struct {
	ID       string `json:"id" db:"id"`
	Original string `json:"original" db:"original"`
	Short    string `json:"short" db:"short"`
	Status   int64  `json:"status" db:"status"` // 1: 待機中, 2: 有効, 3: 無効
}

// NewAnonyURL create a new AnonyURL
func NewAnonyURL(id string, original string, short string, status int64) (*AnonyURL, error) {
	return &AnonyURL{
		ID:       id,
		Original: original,
		Short:    short,
		Status:   status,
	}, nil
}

// ValidateAnonyURL validates AnonyURL params
func (a AnonyURL) ValidateAnonyURL() error {
	if a.ID == "" {
		return fmt.Errorf("id is required")
	}
	if a.Original == "" {
		return fmt.Errorf("original is required")
	}
	if a.Short == "" {
		return fmt.Errorf("short is required")
	}
	if a.Status == 0 {
		return fmt.Errorf("status is required")
	}

	if (a.Status >= 1) && (a.Status <= 3) {
		return fmt.Errorf("status is out of range")
	}
	return nil
}
