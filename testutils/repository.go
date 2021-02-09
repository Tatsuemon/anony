package testutils

import (
	"context"

	"github.com/Tatsuemon/anony/domain/model"
)

// UserRepoMock is mock of userRepository
type UserRepoMock struct {
	FakeFindAll             func() ([]*model.User, error)
	FakeFindByID            func(id string) (*model.User, error)
	FakeFindByName          func(name string) (*model.User, error)
	FakeFindByEmail         func(email string) (*model.User, error)
	FakeFindByNameOrEmail   func(nameOrEmail string) (*model.User, error)
	FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
	FakeSave                func(ctx context.Context, user *model.User) error
	FakeUpdate              func(ctx context.Context, user *model.User) error
	FakeDelete              func(ctx context.Context, user *model.User) error
}

func (m UserRepoMock) FindAll() ([]*model.User, error) {
	return m.FakeFindAll()
}
func (m UserRepoMock) FindByID(id string) (*model.User, error) {
	return m.FakeFindByID(id)
}
func (m UserRepoMock) FindByName(name string) (*model.User, error) {
	return m.FakeFindByName(name)
}
func (m UserRepoMock) FindByEmail(email string) (*model.User, error) {
	return m.FakeFindByEmail(email)
}
func (m UserRepoMock) FindByNameOrEmail(namrOrEmail string) (*model.User, error) {
	return m.FakeFindByNameOrEmail(namrOrEmail)
}
func (m UserRepoMock) FindDuplicatedUsers(name, email string) ([]*model.User, error) {
	return m.FakeFindDuplicatedUsers(name, email)
}
func (m UserRepoMock) Save(ctx context.Context, user *model.User) error {
	return m.FakeSave(ctx, user)
}
func (m UserRepoMock) Update(ctx context.Context, user *model.User) error {
	return m.FakeUpdate(ctx, user)
}
func (m UserRepoMock) Delete(ctx context.Context, user *model.User) error {
	return m.FakeDelete(ctx, user)
}

// AnonyURLRepoMock is mock of anonyURLRepository
type AnonyURLRepoMock struct {
	FakeFindByID               func(id string) (*model.AnonyURL, error)
	FakeFindByUserID           func(userID string) ([]*model.AnonyURL, error)
	FakeFindByUserIDWithStatus func(userID string, status int64) ([]*model.AnonyURL, error)
	FakeFindByOriginalInUser   func(original string, userID string) (*model.AnonyURL, error)
	FakeFindByAnonyURL         func(anonyURL string) (*model.AnonyURL, error)
	FakeGetIDByOriginalUser    func(original, userID string) (string, error)
	FakeSave                   func(ctx context.Context, an *model.AnonyURL, userID string) error
	FakeUpdateStatus           func(ctx context.Context, id string, status int64) error
}

func (a AnonyURLRepoMock) FindByID(id string) (*model.AnonyURL, error) {
	return a.FakeFindByID(id)
}
func (a AnonyURLRepoMock) FindByUserID(userID string) ([]*model.AnonyURL, error) {
	return a.FakeFindByUserID(userID)
}
func (a AnonyURLRepoMock) FindByUserIDWithStatus(userID string, status int64) ([]*model.AnonyURL, error) {
	return a.FakeFindByUserIDWithStatus(userID, status)
}
func (a AnonyURLRepoMock) FindByOriginalInUser(original string, userID string) (*model.AnonyURL, error) {
	return a.FakeFindByOriginalInUser(original, userID)
}
func (a AnonyURLRepoMock) FindByAnonyURL(anonyURL string) (*model.AnonyURL, error) {
	return a.FakeFindByAnonyURL(anonyURL)
}
func (a AnonyURLRepoMock) GetIDByOriginalUser(original, userID string) (string, error) {
	return a.FakeGetIDByOriginalUser(original, userID)
}
func (a AnonyURLRepoMock) Save(ctx context.Context, an *model.AnonyURL, userID string) error {
	return a.FakeSave(ctx, an, userID)
}
func (a AnonyURLRepoMock) UpdateStatus(ctx context.Context, id string, status int64) error {
	return a.FakeUpdateStatus(ctx, id, status)
}
