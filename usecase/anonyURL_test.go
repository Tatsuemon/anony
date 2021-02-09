package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"github.com/Tatsuemon/anony/testutils"
)

func TestNewAnonyURLUseCase(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	tests := []struct {
		name string
		want AnonyURLUseCase
	}{
		{
			name: "NORMAL: 正常にAnonyURLUseCaseが作成できる",
			want: &anonyURLUseCase{
				testutils.AnonyURLRepoMock{},
				transaction,
				testutils.AnonyURLServiceMock{},
			},
		},
	}
	for _, tt := range tests {
		repo := testutils.AnonyURLRepoMock{}
		service := testutils.AnonyURLServiceMock{}
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnonyURLUseCase(repo, transaction, service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnonyURLUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLUseCase_CreateAnonyURL(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "NORMAL: 正常にAnonyURLを作成できる",
			args: args{
				ctx:    context.Background(),
				userID: "abcdefghijklmnopqrstuvwxyz1234567890",
			},
			want:    "http://localhost-test/z1234567/",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &anonyURLUseCase{
				repo:        testutils.AnonyURLRepoMock{},
				transaction: transaction,
				service:     testutils.AnonyURLServiceMock{},
			}
			got, err := u.CreateAnonyURL(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLUseCase.CreateAnonyURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if len(got) != 39 {
					t.Errorf("anonyURLUseCase.CreateAnonyURL() len(got) = %v, want 39", len(got))
				}
			}
			// 最後の部分はランダムであるため, そこまでを検証する
			if got[:31] != tt.want[:31] {
				t.Errorf("anonyURLUseCase.CreateAnonyURL() = %v, want %v", got[:31], tt.want[:31])
			}
		})
	}
}

func Test_anonyURLUseCase_SaveAnonyURL(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByID            func(id string) (*model.AnonyURL, error)
		FakeGetIDByOriginalUser func(original, userID string) (string, error)
		FakeSave                func(ctx context.Context, an *model.AnonyURL, userID string) error
		FakeUpdateStatus        func(ctx context.Context, id string, status int64) error
	}
	type serviceMocks struct {
		FakeExistID             func(id string) (bool, error)
		FakeExistOriginalInUser func(original, userID string) (bool, error)
	}
	type args struct {
		ctx    context.Context
		an     *model.AnonyURL
		userID string
	}
	tests := []struct {
		name         string
		args         args
		repoMocks    repoMocks
		serviceMocks serviceMocks
		want         *model.AnonyURL
		wantErr      bool
	}{
		{
			name: "NORMAL: 新規作成",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id1",
						Original: "http://localhost:8888/original1",
						Short:    "http://localhost:8888/short1",
						Status:   1,
					}, nil
				},
				FakeSave: func(ctx context.Context, an *model.AnonyURL, userID string) error {
					return nil
				},
			},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return false, nil
				},
			},
			want: &model.AnonyURL{
				ID:       "id1",
				Original: "http://localhost:8888/original1",
				Short:    "http://localhost:8888/short1",
				Status:   1,
			},
			wantErr: false,
		},
		{
			name: "NORMAL: 既にある場合",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id1",
						Original: "http://localhost:8888/original1",
						Short:    "http://localhost:8888/short1",
						Status:   1,
					}, nil
				},
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "id1", nil
				},
				FakeUpdateStatus: func(ctx context.Context, id string, status int64) error {
					return nil
				},
			},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return true, nil
				},
			},
			want: &model.AnonyURL{
				ID:       "id1",
				Original: "http://localhost:8888/original1",
				Short:    "http://localhost:8888/short1",
				Status:   1,
			},
			wantErr: false,
		},
		{
			name: "ERROR: service.ExistOriginalInUserでErrorを返す",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{},
			serviceMocks: serviceMocks{
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return false, fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: service.ExistIDでErrorを返す",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, fmt.Errorf("error")
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return false, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: IDが既に存在している場合",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return true, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return false, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: AnonyURLのバリデーションに引っかかる場合",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return false, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: 新規作成でSaveがErrorを返す",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id1",
						Original: "http://localhost:8888/original1",
						Short:    "http://localhost:8888/short1",
						Status:   1,
					}, nil
				},
				FakeSave: func(ctx context.Context, an *model.AnonyURL, userID string) error {
					return fmt.Errorf("error")
				},
			},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return false, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: 既にある場合で, repo.GetIDByOriginalUserがErrorを返す場合",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "", fmt.Errorf("error")
				},
			},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return true, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: 既にある場合で, repo.UpdateStatusがErrorを返す場合",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id1",
						Original: "http://localhost:8888/original1",
						Short:    "http://localhost:8888/short1",
						Status:   1,
					}, nil
				},
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "id1", nil
				},
				FakeUpdateStatus: func(ctx context.Context, id string, status int64) error {
					return fmt.Errorf("error")
				},
			},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return true, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: repo.FindByIDがErrorを返す場合",
			args: args{
				ctx: context.Background(),
				an: &model.AnonyURL{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				userID: "user_id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "id1", nil
				},
				FakeUpdateStatus: func(ctx context.Context, id string, status int64) error {
					return nil
				},
			},
			serviceMocks: serviceMocks{
				FakeExistID: func(id string) (bool, error) {
					return false, nil
				},
				FakeExistOriginalInUser: func(original, userID string) (bool, error) {
					return true, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.AnonyURLRepoMock{
				FakeFindByID:            tt.repoMocks.FakeFindByID,
				FakeGetIDByOriginalUser: tt.repoMocks.FakeGetIDByOriginalUser,
				FakeSave:                tt.repoMocks.FakeSave,
				FakeUpdateStatus:        tt.repoMocks.FakeUpdateStatus,
			}
			service := testutils.AnonyURLServiceMock{
				FakeExistID:             tt.serviceMocks.FakeExistID,
				FakeExistOriginalInUser: tt.serviceMocks.FakeExistOriginalInUser,
			}
			u := &anonyURLUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.SaveAnonyURL(tt.args.ctx, tt.args.an, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLUseCase.SaveAnonyURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anonyURLUseCase.SaveAnonyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLUseCase_UpdateAnonyURLStatus(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByID            func(id string) (*model.AnonyURL, error)
		FakeGetIDByOriginalUser func(original, userID string) (string, error)
		FakeUpdateStatus        func(ctx context.Context, id string, status int64) error
	}
	type args struct {
		ctx      context.Context
		original string
		userID   string
		status   int64
	}
	tests := []struct {
		name      string
		args      args
		repoMocks repoMocks
		want      *model.AnonyURL
		wantErr   bool
	}{
		{
			name: "NORMAL: Statusの更新ができる",
			args: args{
				ctx:      context.Background(),
				original: "http://localhost:8888/original1",
				userID:   "user_id",
				status:   1,
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id1",
						Original: "http://localhost:8888/original1",
						Short:    "http://localhost:8888/short1",
						Status:   1,
					}, nil
				},
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "id1", nil
				},
				FakeUpdateStatus: func(ctx context.Context, id string, status int64) error {
					return nil
				},
			},
			want: &model.AnonyURL{
				ID:       "id1",
				Original: "http://localhost:8888/original1",
				Short:    "http://localhost:8888/short1",
				Status:   1,
			},
			wantErr: false,
		},
		{
			name: "ERROR: repo.GetIDByOriginalUserでErrorを返す",
			args: args{
				ctx:      context.Background(),
				original: "http://localhost:8888/original1",
				userID:   "user_id",
				status:   1,
			},
			repoMocks: repoMocks{
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "", fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: repo.GetIDByOriginalUserが空文字を返す",
			args: args{
				ctx:      context.Background(),
				original: "http://localhost:8888/original1",
				userID:   "user_id",
				status:   1,
			},
			repoMocks: repoMocks{
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "", nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: repo.UpdateStatusがErrorを返す",
			args: args{
				ctx:      context.Background(),
				original: "http://localhost:8888/original1",
				userID:   "user_id",
				status:   1,
			},
			repoMocks: repoMocks{
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "id1", nil
				},
				FakeUpdateStatus: func(ctx context.Context, id string, status int64) error {
					return fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: repo.FindByIDがErrorを返す",
			args: args{
				ctx:      context.Background(),
				original: "http://localhost:8888/original1",
				userID:   "user_id",
				status:   1,
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
				FakeGetIDByOriginalUser: func(original, userID string) (string, error) {
					return "id1", nil
				},
				FakeUpdateStatus: func(ctx context.Context, id string, status int64) error {
					return nil
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.AnonyURLRepoMock{
				FakeFindByID:            tt.repoMocks.FakeFindByID,
				FakeGetIDByOriginalUser: tt.repoMocks.FakeGetIDByOriginalUser,
				FakeUpdateStatus:        tt.repoMocks.FakeUpdateStatus,
			}
			service := testutils.AnonyURLServiceMock{}
			u := &anonyURLUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.UpdateAnonyURLStatus(tt.args.ctx, tt.args.original, tt.args.userID, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLUseCase.UpdateAnonyURLStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anonyURLUseCase.UpdateAnonyURLStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLUseCase_ListAnonyURLs(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByUserID           func(userID string) ([]*model.AnonyURL, error)
		FakeFindByUserIDWithStatus func(userID string, status int64) ([]*model.AnonyURL, error)
	}
	type args struct {
		ctx    context.Context
		userID string
		q      int64
	}
	tests := []struct {
		name      string
		args      args
		repoMocks repoMocks
		want      []*model.AnonyURL
		wantErr   bool
	}{
		{
			name: "NORMAL: q=0の場合",
			args: args{
				ctx:    context.Background(),
				userID: "user_id",
				q:      0,
			},
			repoMocks: repoMocks{
				FakeFindByUserID: func(userID string) ([]*model.AnonyURL, error) {
					return []*model.AnonyURL{
						{
							ID:       "id1",
							Original: "http://localhost:8888/original1",
							Short:    "http://localhost:8888/short1",
							Status:   1,
						},
						{
							ID:       "id2",
							Original: "http://localhost:8888/original2",
							Short:    "http://localhost:8888/short2",
							Status:   2,
						},
					}, nil
				},
			},
			want: []*model.AnonyURL{
				{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				{
					ID:       "id2",
					Original: "http://localhost:8888/original2",
					Short:    "http://localhost:8888/short2",
					Status:   2,
				},
			},
			wantErr: false,
		},
		{
			name: "NORMAL: q=1の場合",
			args: args{
				ctx:    context.Background(),
				userID: "user_id",
				q:      1,
			},
			repoMocks: repoMocks{
				FakeFindByUserIDWithStatus: func(userID string, status int64) ([]*model.AnonyURL, error) {
					return []*model.AnonyURL{
						{
							ID:       "id1",
							Original: "http://localhost:8888/original1",
							Short:    "http://localhost:8888/short1",
							Status:   1,
						},
						{
							ID:       "id2",
							Original: "http://localhost:8888/original2",
							Short:    "http://localhost:8888/short2",
							Status:   1,
						},
					}, nil
				},
			},
			want: []*model.AnonyURL{
				{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   1,
				},
				{
					ID:       "id2",
					Original: "http://localhost:8888/original2",
					Short:    "http://localhost:8888/short2",
					Status:   1,
				},
			},
			wantErr: false,
		},
		{
			name: "NORMAL: q=2の場合",
			args: args{
				ctx:    context.Background(),
				userID: "user_id",
				q:      2,
			},
			repoMocks: repoMocks{
				FakeFindByUserIDWithStatus: func(userID string, status int64) ([]*model.AnonyURL, error) {
					return []*model.AnonyURL{
						{
							ID:       "id1",
							Original: "http://localhost:8888/original1",
							Short:    "http://localhost:8888/short1",
							Status:   2,
						},
						{
							ID:       "id2",
							Original: "http://localhost:8888/original2",
							Short:    "http://localhost:8888/short2",
							Status:   2,
						},
					}, nil
				},
			},
			want: []*model.AnonyURL{
				{
					ID:       "id1",
					Original: "http://localhost:8888/original1",
					Short:    "http://localhost:8888/short1",
					Status:   2,
				},
				{
					ID:       "id2",
					Original: "http://localhost:8888/original2",
					Short:    "http://localhost:8888/short2",
					Status:   2,
				},
			},
			wantErr: false,
		},
		{
			name: "ERROR: q!=0, 1, の場合",
			args: args{
				ctx:    context.Background(),
				userID: "user_id",
				q:      3,
			},
			repoMocks: repoMocks{},
			want:      nil,
			wantErr:   true,
		},
		{
			name: "ERROR: q=0の場合にFindByUserIDがErrorを返す",
			args: args{
				ctx:    context.Background(),
				userID: "user_id",
				q:      0,
			},
			repoMocks: repoMocks{
				FakeFindByUserID: func(userID string) ([]*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: q=1の場合にFindByUserIDがErrorを返す",
			args: args{
				ctx:    context.Background(),
				userID: "user_id",
				q:      1,
			},
			repoMocks: repoMocks{
				FakeFindByUserIDWithStatus: func(userID string, status int64) ([]*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: q=2の場合にFindByUserIDがErrorを返す",
			args: args{
				ctx:    context.Background(),
				userID: "user_id",
				q:      2,
			},
			repoMocks: repoMocks{
				FakeFindByUserIDWithStatus: func(userID string, status int64) ([]*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.AnonyURLRepoMock{
				FakeFindByUserID:           tt.repoMocks.FakeFindByUserID,
				FakeFindByUserIDWithStatus: tt.repoMocks.FakeFindByUserIDWithStatus,
			}
			service := testutils.AnonyURLServiceMock{}
			u := &anonyURLUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.ListAnonyURLs(tt.args.ctx, tt.args.userID, tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLUseCase.ListAnonyURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anonyURLUseCase.ListAnonyURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anonyURLUseCase_GetOriginalByAnonyURL(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByAnonyURL func(anonyURL string) (*model.AnonyURL, error)
	}
	type args struct {
		ctx      context.Context
		anonyURL string
	}
	tests := []struct {
		name      string
		args      args
		repoMocks repoMocks
		want      string
		wantErr   bool
	}{
		{
			name: "NORMAL: Original URLを返す",
			args: args{
				ctx:      context.Background(),
				anonyURL: "http://localhost:8888/aaaabbbb/ccccdddd",
			},
			repoMocks: repoMocks{
				FakeFindByAnonyURL: func(anonyURL string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id",
						Original: "http://localhost:8888/original",
						Short:    "http://localhost:8888/aaaabbbb/ccccdddd",
						Status:   1,
					}, nil
				},
			},
			want:    "http://localhost:8888/original",
			wantErr: false,
		},
		{
			name: "NORMAL: Statusが1以外の時は空文字を返す",
			args: args{
				ctx:      context.Background(),
				anonyURL: "http://localhost:8888/aaaabbbb/ccccdddd",
			},
			repoMocks: repoMocks{
				FakeFindByAnonyURL: func(anonyURL string) (*model.AnonyURL, error) {
					return &model.AnonyURL{
						ID:       "id",
						Original: "http://localhost:8888/original",
						Short:    "http://localhost:8888/aaaabbbb/ccccdddd",
						Status:   0,
					}, nil
				},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "NORMAL: AnonyURLが見つからなかった場合は空文字を返す",
			args: args{
				ctx:      context.Background(),
				anonyURL: "http://localhost:8888/aaaabbbb/ccccdddd",
			},
			repoMocks: repoMocks{
				FakeFindByAnonyURL: func(anonyURL string) (*model.AnonyURL, error) {
					return nil, nil
				},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "ERROR: FindByAnonyURLでErrorを返す場合",
			args: args{
				ctx:      context.Background(),
				anonyURL: "http://localhost:8888/aaaabbbb/ccccdddd",
			},
			repoMocks: repoMocks{
				FakeFindByAnonyURL: func(anonyURL string) (*model.AnonyURL, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.AnonyURLRepoMock{
				FakeFindByAnonyURL: tt.repoMocks.FakeFindByAnonyURL,
			}
			service := testutils.AnonyURLServiceMock{}
			u := &anonyURLUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.GetOriginalByAnonyURL(tt.args.ctx, tt.args.anonyURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("anonyURLUseCase.GetOriginalByAnonyURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("anonyURLUseCase.GetOriginalByAnonyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
