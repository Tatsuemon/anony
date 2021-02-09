package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"github.com/Tatsuemon/anony/testutils"
	"github.com/Tatsuemon/anony/usecase/dto"
)

func TestNewAnonyURLWithUserUseCase(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	tests := []struct {
		name string
		want AnonyURLWithUserUseCase
	}{
		{
			name: "NORMAL: AnonyURLWithUserUseCaseの作成ができる",
			want: &anonyURLWithUserUseCase{
				testutils.UserAnonyURLAccessorMock{},
				transaction,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := testutils.UserAnonyURLAccessorMock{}
			if got := NewAnonyURLWithUserUseCase(a, transaction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnonyURLWithUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLWithUserUseCase_CountByUser(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type accMocks struct {
		FakeCountAnonyURLByUser func(userID string) (*dto.AnonyURLCountByUser, error)
	}
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name     string
		args     args
		accMocks accMocks
		want     *dto.AnonyURLCountByUser
		wantErr  bool
	}{
		{
			name: "NORMAL: 正常にusecase.CoutAnonyURLByUserが値を返す",
			args: args{
				ctx:    context.Background(),
				userID: "id",
			},
			accMocks: accMocks{
				FakeCountAnonyURLByUser: func(userID string) (*dto.AnonyURLCountByUser, error) {
					return &dto.AnonyURLCountByUser{
						Name:          "name",
						Email:         "email",
						CntURLs:       10,
						CntActiveURLs: 5,
					}, nil
				},
			},
			want: &dto.AnonyURLCountByUser{
				Name:          "name",
				Email:         "email",
				CntURLs:       10,
				CntActiveURLs: 5,
			},
			wantErr: false,
		},
		{
			name: "ERROR: usecase.CoutAnonyURLByUserがERRORを返す",
			args: args{
				ctx:    context.Background(),
				userID: "id",
			},
			accMocks: accMocks{
				FakeCountAnonyURLByUser: func(userID string) (*dto.AnonyURLCountByUser, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := anonyURLWithUserUseCase{
				accessor: testutils.UserAnonyURLAccessorMock{
					FakeCountAnonyURLByUser: tt.accMocks.FakeCountAnonyURLByUser,
				},
				transaction: transaction,
			}
			got, err := a.CountByUser(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLWithUserUseCase.CountByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anonyURLWithUserUseCase.CountByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
