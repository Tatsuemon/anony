package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/Tatsuemon/anony/domain/model"
)

type userRepositoryMock struct {
	FakeFindAll             func() ([]*model.User, error)
	FakeFindByID            func(id string) (*model.User, error)
	FakeFindByName          func(name string) (*model.User, error)
	FakeFindByEmail         func(email string) (*model.User, error)
	FakeFindByNameOrEmail   func(nameOrEmail string) (*model.User, error)
	FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
	FakeStore               func(ctx context.Context, user *model.User) (*model.User, error)
	FakeUpdate              func(ctx context.Context, user *model.User) (*model.User, error)
	FakeDelete              func(ctx context.Context, user *model.User) error
}

func (m userRepositoryMock) FindAll() ([]*model.User, error) {
	return m.FakeFindAll()
}
func (m userRepositoryMock) FindByID(id string) (*model.User, error) {
	return m.FakeFindByID(id)
}
func (m userRepositoryMock) FindByName(name string) (*model.User, error) {
	return m.FakeFindByName(name)
}
func (m userRepositoryMock) FindByEmail(email string) (*model.User, error) {
	return m.FakeFindByEmail(email)
}
func (m userRepositoryMock) FindByNameOrEmail(namrOrEmail string) (*model.User, error) {
	return m.FakeFindByNameOrEmail(namrOrEmail)
}
func (m userRepositoryMock) FindDuplicatedUsers(name, email string) ([]*model.User, error) {
	return m.FakeFindDuplicatedUsers(name, email)
}
func (m userRepositoryMock) Store(ctx context.Context, user *model.User) (*model.User, error) {
	return m.FakeStore(ctx, user)
}
func (m userRepositoryMock) Update(ctx context.Context, user *model.User) (*model.User, error) {
	return m.FakeUpdate(ctx, user)
}
func (m userRepositoryMock) Delete(ctx context.Context, user *model.User) error {
	return m.FakeDelete(ctx, user)
}

func TestNewUserService(t *testing.T) {
	type mocks struct {
		// FakeFindAll             func() ([]*model.User, error)
		FakeFindByID func(id string) (*model.User, error)
		// FakeFindByName          func(name string) (*model.User, error)
		// FakeFindByEmail         func(email string) (*model.User, error)
		// FakeFindByNameOrEmail   func(nameOrEmail string) (*model.User, error)
		// FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
		// FakeStore               func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeUpdate              func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeDelete              func(ctx context.Context, user *model.User) error
	}
	tests := []struct {
		name  string
		mocks mocks
		want  UserService
	}{
		{
			name: "NORMAL: ",
			mocks: mocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "name",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
			},
			want: &userService{},
		},
	}
	for _, tt := range tests {
		repo := userRepositoryMock{
			FakeFindByID: tt.mocks.FakeFindByID,
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(repo); got == nil {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_ExistsID(t *testing.T) {
	type mocks struct {
		// FakeFindAll             func() ([]*model.User, error)
		FakeFindByID func(id string) (*model.User, error)
		// FakeFindByName          func(name string) (*model.User, error)
		// FakeFindByEmail         func(email string) (*model.User, error)
		// FakeFindByNameOrEmail   func(nameOrEmail string) (*model.User, error)
		// FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
		// FakeStore               func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeUpdate              func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeDelete              func(ctx context.Context, user *model.User) error
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    bool
		wantErr bool
	}{
		{
			name: "NORMAL: userRepository.FindByIDで値を返す",
			args: args{
				id: "id",
			},
			mocks: mocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "name",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ERROR: userRepository.FindByIDでエラーを返す",
			args: args{
				id: "id",
			},
			mocks: mocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "name",
						Email:         "email",
						EncryptedPass: "password",
					}, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "ERROR: userRepository.FindByIDでエラーを返す userがnil",
			args: args{
				id: "id",
			},
			mocks: mocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				UserRepository: userRepositoryMock{
					FakeFindByID: tt.mocks.FakeFindByID,
				},
			}
			got, err := u.ExistsID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.ExistsID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.ExistsID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_ExistsName(t *testing.T) {
	type mocks struct {
		// FakeFindAll             func() ([]*model.User, error)
		// FakeFindByID            func(id string) (*model.User, error)
		FakeFindByName func(name string) (*model.User, error)
		// FakeFindByEmail         func(email string) (*model.User, error)
		// FakeFindByNameOrEmail   func(nameOrEmail string) (*model.User, error)
		// FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
		// FakeStore               func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeUpdate              func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeDelete              func(ctx context.Context, user *model.User) error
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    bool
		wantErr bool
	}{
		{
			name: "NORMAL: userRepository.FindByNameで値を返す場合",
			args: args{
				name: "name",
			},
			mocks: mocks{
				FakeFindByName: func(name string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "name",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ERROR: userRepository.FindByNameでErrorを返す場合",
			args: args{
				name: "name",
			},
			mocks: mocks{
				FakeFindByName: func(name string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "name",
						Email:         "email",
						EncryptedPass: "password",
					}, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "ERROR: userRepository.FindByNameでErrorを返す場合",
			args: args{
				name: "name",
			},
			mocks: mocks{
				FakeFindByName: func(name string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				UserRepository: userRepositoryMock{
					FakeFindByName: tt.mocks.FakeFindByName,
				},
			}
			got, err := u.ExistsName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.ExistsName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.ExistsName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_ExistsEmail(t *testing.T) {
	type mocks struct {
		// FakeFindAll             func() ([]*model.User, error)
		// FakeFindByID            func(id string) (*model.User, error)
		// FakeFindByName          func(name string) (*model.User, error)
		FakeFindByEmail func(email string) (*model.User, error)
		// FakeFindByNameOrEmail   func(nameOrEmail string) (*model.User, error)
		// FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
		// FakeStore               func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeUpdate              func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeDelete              func(ctx context.Context, user *model.User) error
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    bool
		wantErr bool
	}{
		{
			name: "NORMAL: userRepository.FindByEmailで値を返す場合",
			args: args{
				email: "email",
			},
			mocks: mocks{
				FakeFindByEmail: func(email string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "name",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ERROR: userRepository.FindByEmailでErrorを返す場合",
			args: args{
				email: "email",
			},
			mocks: mocks{
				FakeFindByEmail: func(email string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "ERROR: userRepository.FindByEmailでErrorを返す場合",
			args: args{
				email: "email",
			},
			mocks: mocks{
				FakeFindByEmail: func(email string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "name",
						Email:         "email",
						EncryptedPass: "password",
					}, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				UserRepository: userRepositoryMock{
					FakeFindByEmail: tt.mocks.FakeFindByEmail,
				},
			}
			got, err := u.ExistsEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.ExistsEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.ExistsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_ExistsDuplicatedUser(t *testing.T) {
	type mocks struct {
		// FakeFindAll             func() ([]*model.User, error)
		// FakeFindByID            func(id string) (*model.User, error)
		// FakeFindByName          func(name string) (*model.User, error)
		// FakeFindByEmail         func(email string) (*model.User, error)
		// FakeFindByNameOrEmail   func(nameOrEmail string) (*model.User, error)
		FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
		// FakeStore               func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeUpdate              func(ctx context.Context, user *model.User) (*model.User, error)
		// FakeDelete              func(ctx context.Context, user *model.User) error
	}
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    bool
		wantErr bool
	}{
		{
			name: "NORMAL: userRepository.FindByDuplicatdUsersでUserの配列を返す場合",
			args: args{
				name:  "name",
				email: "email",
			},
			mocks: mocks{
				FakeFindDuplicatedUsers: func(name, email string) ([]*model.User, error) {
					return []*model.User{
						{
							ID:            "id",
							Name:          "name",
							Email:         "email",
							EncryptedPass: "password",
						},
					}, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "NORMAL: userRepository.FindByDuplicatdUsersでUserの配列が空の場合",
			args: args{
				name:  "name",
				email: "email",
			},
			mocks: mocks{
				FakeFindDuplicatedUsers: func(name, email string) ([]*model.User, error) {
					return []*model.User{}, nil
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "ERROR: userRepository.FindByDuplicatdUsersでErrorを返す場合",
			args: args{
				name:  "name",
				email: "email",
			},
			mocks: mocks{
				FakeFindDuplicatedUsers: func(name, email string) ([]*model.User, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				UserRepository: userRepositoryMock{
					FakeFindDuplicatedUsers: tt.mocks.FakeFindDuplicatedUsers,
				},
			}
			got, err := u.ExistsDuplicatedUser(tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.ExistsDuplicatedUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.ExistsDuplicatedUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
