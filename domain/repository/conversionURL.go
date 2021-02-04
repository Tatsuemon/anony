package repository

import (
	"github.com/Tatsuemon/anony/domain/model"
)

// ConversionURLRepository is a interface
type ConversionURLRepository interface {
	FindByID(id string) ([]*model.ConversionURL, error)
	FindByOriginalURLOfUser(original string, userID string) ([]*model.ConversionURL, error)
}
