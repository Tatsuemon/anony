package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Tatsuemon/anony/domain/model"
)

type anonyURLRepoMock struct {
	FakeFindByID             func(id string) (*model.AnonyURL, error)
	FakeFindByOriginalInUser func(original string, userID string) (*model.AnonyURL, error)
	FakeSave                 func(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error)
}

func (a anonyURLRepoMock) FindByID(id string) (*model.AnonyURL, error) {
	return a.FakeFindByID(id)
}
func (a anonyURLRepoMock) FindByOriginalInUser(original string, userID string) (*model.AnonyURL, error) {
	return a.FakeFindByOriginalInUser(original, userID)
}
func (a anonyURLRepoMock) Save(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error) {
	return a.FakeSave(ctx, an, userID)
}

func TestNewAnonyURLService(t *testing.T) {
	tests := []struct {
		name string
		want AnonyURLService
	}{
		{
			name: "NORMAL: NewAnonyURLService",
			want: &anonyURLService{anonyURLRepoMock{}},
		},
	}
	for _, tt := range tests {
		repo := anonyURLRepoMock{}
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnonyURLService(repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnonyURLService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLService_ExistID(t *testing.T) {
	type mocks struct {
		FakeFindByID             func(id string) (*model.AnonyURL, error)
		FakeFindByOriginalInUser func(original string, userID string) (*model.AnonyURL, error)
		FakeSave                 func(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error)
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
			name: "NORMAL: 重複するものが存在しない場合",
			args: args{
				id: "id",
			},
			mocks: mocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return nil, nil
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "NORMAL: 重複するものが存在する場合",
			args: args{
				id: "id",
			},
			mocks: mocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id",
						Original: "original",
						Short:    "short",
						Status:   1,
					}, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ERROR: anonyURLRepo.FindByIDがErrorを返す",
			args: args{
				id: "id",
			},
			mocks: mocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &anonyURLService{
				repo: anonyURLRepoMock{
					FakeFindByID: tt.mocks.FakeFindByID,
				},
			}
			got, err := a.ExistID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLService.ExistID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("anonyURLService.ExistID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLService_ExistOriginalInUser(t *testing.T) {
	type mocks struct {
		FakeFindByID             func(id string) (*model.AnonyURL, error)
		FakeFindByOriginalInUser func(original string, userID string) (*model.AnonyURL, error)
		FakeSave                 func(ctx context.Context, an *model.AnonyURL, userID string) (*model.AnonyURL, error)
	}
	type args struct {
		original string
		userID   string
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    bool
		wantErr bool
	}{
		{
			name: "NORMAL: 重複するものが存在しない場合",
			args: args{
				original: "original",
				userID:   "user-id",
			},
			mocks: mocks{
				FakeFindByOriginalInUser: func(original string, userID string) (*model.AnonyURL, error) {
					return nil, nil
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "NORMAL: 重複するものが存在する場合",
			args: args{
				original: "original",
				userID:   "user-id",
			},
			mocks: mocks{
				FakeFindByOriginalInUser: func(original string, userID string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id",
						Original: "original",
						Short:    "short",
						Status:   1,
					}, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ERROR: anonyURLRepo.FindByOriginalInUserがERRORを返す時",
			args: args{
				original: "original",
				userID:   "user-id",
			},
			mocks: mocks{
				FakeFindByOriginalInUser: func(original string, userID string) (*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &anonyURLService{
				repo: anonyURLRepoMock{
					FakeFindByOriginalInUser: tt.mocks.FakeFindByOriginalInUser,
				},
			}
			got, err := a.ExistOriginalInUser(tt.args.original, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLService.ExistOriginalInUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("anonyURLService.ExistOriginalInUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
