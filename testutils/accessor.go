package testutils

import "github.com/Tatsuemon/anony/usecase/dto"

// UserAnonyURLAccessorMock is mock of UserAnonyURLAccessor
type UserAnonyURLAccessorMock struct {
	FakeCountAnonyURLByUser func(userID string) (*dto.AnonyURLCountByUser, error)
}

func (m UserAnonyURLAccessorMock) CountAnonyURLByUser(userID string) (*dto.AnonyURLCountByUser, error) {
	return m.FakeCountAnonyURLByUser(userID)
}
