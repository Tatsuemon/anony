package queryservice

import (
	"github.com/Tatsuemon/anony/usecase/dto"
)

type UserAnonyURLAccessor interface {
	CountAnonyURLByUser(userID string) (*dto.AnonyURLCountByUser, error)
}
