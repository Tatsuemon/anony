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
