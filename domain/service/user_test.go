package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/testutils"
)

func TestNewUserService(t *testing.T) {
	tests := []struct {
		name string
		want UserService
	}{
		{
			name: "NORMAL: NewUserService",
			want: &userService{testutils.UserRepoMock{}},
		},
	}
	for _, tt := range tests {
		repo := testutils.UserRepoMock{}
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_ExistsID(t *testing.T) {
	type mocks struct {
		FakeFindByID func(id string) (*model.User, error)
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
			name: "NORMAL: userがnilの場合",
			args: args{
				id: "id",
			},
			mocks: mocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return nil, nil
				},
			},
			want:    false,
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
				UserRepository: testutils.UserRepoMock{
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
		FakeFindByName func(name string) (*model.User, error)
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
			name: "NORMAL: userがnilの場合",
			args: args{
				name: "name",
			},
			mocks: mocks{
				FakeFindByName: func(name string) (*model.User, error) {
					return nil, nil
				},
			},
			want:    false,
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
				UserRepository: testutils.UserRepoMock{
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
		FakeFindByEmail func(email string) (*model.User, error)
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
			name: "NORMAL: userがnilの場合",
			args: args{
				email: "email",
			},
			mocks: mocks{
				FakeFindByEmail: func(email string) (*model.User, error) {
					return nil, nil
				},
			},
			want:    false,
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
				UserRepository: testutils.UserRepoMock{
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
		FakeFindDuplicatedUsers func(name, email string) ([]*model.User, error)
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
				UserRepository: testutils.UserRepoMock{
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
