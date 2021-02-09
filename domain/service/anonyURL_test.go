package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/testutils"
)

func TestNewAnonyURLService(t *testing.T) {
	tests := []struct {
		name string
		want AnonyURLService
	}{
		{
			name: "NORMAL: NewAnonyURLService",
			want: &anonyURLService{testutils.AnonyURLRepoMock{}},
		},
	}
	for _, tt := range tests {
		repo := testutils.AnonyURLRepoMock{}
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnonyURLService(repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnonyURLService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLService_ExistID(t *testing.T) {
	type mocks struct {
		FakeFindByID func(id string) (*model.AnonyURL, error)
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
				repo: testutils.AnonyURLRepoMock{
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
		FakeFindByOriginalInUser func(original string, userID string) (*model.AnonyURL, error)
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
				repo: testutils.AnonyURLRepoMock{
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
