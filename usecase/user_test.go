package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"github.com/Tatsuemon/anony/testutils"
)

func TestNewUserUseCase(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	tests := []struct {
		name string
		want UserUseCase
	}{
		{
			name: "NORMAL: userUseCaseの作成ができる",
			want: &userUseCase{
				testutils.UserRepoMock{},
				transaction,
				testutils.UserServiceMock{},
			},
		},
	}
	for _, tt := range tests {
		repo := testutils.UserRepoMock{}
		service := testutils.UserServiceMock{}
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUseCase(repo, transaction, service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCase_CreateUser(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByID func(id string) (*model.User, error)
		FakeSave     func(ctx context.Context, user *model.User) error
	}
	type serviceMocks struct {
		FakeExistsID func(id string) (bool, error)
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name         string
		args         args
		repoMocks    repoMocks
		serviceMocks serviceMocks
		want         *model.User
		wantErr      bool
	}{
		{
			name: "NORMAL: 正常に登録できる",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "user",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
				FakeSave: func(ctx context.Context, user *model.User) error {
					return nil
				},
			},
			serviceMocks: serviceMocks{
				FakeExistsID: func(id string) (bool, error) {
					return false, nil
				},
			},
			want: &model.User{
				ID:            "id",
				Name:          "user",
				Email:         "email",
				EncryptedPass: "password",
			},
			wantErr: false,
		},
		{
			name: "ERROR: IDが既に存在する場合",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{},
			serviceMocks: serviceMocks{
				FakeExistsID: func(id string) (bool, error) {
					return true, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: IDが重複でErrorが出た場合",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{},
			serviceMocks: serviceMocks{
				FakeExistsID: func(id string) (bool, error) {
					return false, fmt.Errorf("Error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: Userのバリデーションに引っかかる",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{},
			serviceMocks: serviceMocks{
				FakeExistsID: func(id string) (bool, error) {
					return false, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: repo.SaveがErrorを返す",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{
				FakeSave: func(ctx context.Context, user *model.User) error {
					return fmt.Errorf("error")
				},
			},
			serviceMocks: serviceMocks{
				FakeExistsID: func(id string) (bool, error) {
					return false, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: repo.FindByIDがErrorを返す",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
				FakeSave: func(ctx context.Context, user *model.User) error {
					return nil
				},
			},
			serviceMocks: serviceMocks{
				FakeExistsID: func(id string) (bool, error) {
					return false, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.UserRepoMock{
				FakeFindByID: tt.repoMocks.FakeFindByID,
				FakeSave:     tt.repoMocks.FakeSave,
			}
			service := testutils.UserServiceMock{
				FakeExistsID: tt.serviceMocks.FakeExistsID,
			}
			u := &userUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCase_CheckDuplicatedUser(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type serviceMocks struct {
		FakeExistsDuplicatedUser func(name, email string) (bool, error)
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name         string
		args         args
		serviceMocks serviceMocks
		want         bool
		wantErr      bool
	}{
		{
			name: "NORMAL: nameとemailが重複したユーザーが存在しない",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "name",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			serviceMocks: serviceMocks{
				FakeExistsDuplicatedUser: func(name, email string) (bool, error) {
					return false, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "NORMAL: nameとemailが重複したユーザーが存在する",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "name",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			serviceMocks: serviceMocks{
				FakeExistsDuplicatedUser: func(name, email string) (bool, error) {
					return true, nil
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "ERROR: service.ExistsDuplicatedUserがErrorを返す",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "name",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			serviceMocks: serviceMocks{
				FakeExistsDuplicatedUser: func(name, email string) (bool, error) {
					return false, fmt.Errorf("error")
				},
			},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.UserRepoMock{}
			service := testutils.UserServiceMock{
				FakeExistsDuplicatedUser: tt.serviceMocks.FakeExistsDuplicatedUser,
			}
			u := &userUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.CheckDuplicatedUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.CheckDuplicatedUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUseCase.CheckDuplicatedUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCase_VerifyByNameOrEmailPass(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByNameOrEmail func(nameOrEmail string) (*model.User, error)
	}
	type args struct {
		ctx         context.Context
		nameOrEmail string
		password    string
	}
	pass, _ := model.EncryptPassword("password")
	tests := []struct {
		name      string
		args      args
		repoMocks repoMocks
		want      *model.User
		wantErr   bool
	}{
		{
			name: "NORMAL: 正常にバリデーションを行う",
			args: args{
				ctx:         context.Background(),
				nameOrEmail: "name",
				password:    "password",
			},
			repoMocks: repoMocks{
				FakeFindByNameOrEmail: func(nameOrEmail string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "user",
						Email:         "email",
						EncryptedPass: pass,
					}, nil
				},
			},
			want: &model.User{
				ID:            "id",
				Name:          "user",
				Email:         "email",
				EncryptedPass: pass,
			},
			wantErr: false,
		},
		{
			name: "ERROR: repo.FindByNameOrEmailPassでErrorが出る, ユーザーが見つからないも入る",
			args: args{
				ctx:         context.Background(),
				nameOrEmail: "name",
				password:    "password",
			},
			repoMocks: repoMocks{
				FakeFindByNameOrEmail: func(nameOrEmail string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: パスワードが一致しない",
			args: args{
				ctx:         context.Background(),
				nameOrEmail: "name",
				password:    "password1",
			},
			repoMocks: repoMocks{
				FakeFindByNameOrEmail: func(nameOrEmail string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "user",
						Email:         "email",
						EncryptedPass: pass,
					}, nil
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.UserRepoMock{
				FakeFindByNameOrEmail: tt.repoMocks.FakeFindByNameOrEmail,
			}
			service := testutils.UserServiceMock{}
			u := &userUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.VerifyByNameOrEmailPass(tt.args.ctx, tt.args.nameOrEmail, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.VerifyByNameOrEmailPass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCase.VerifyByNameOrEmailPass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCase_UpdateUser(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByID func(id string) (*model.User, error)
		FakeUpdate   func(ctx context.Context, user *model.User) error
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name      string
		args      args
		repoMocks repoMocks
		want      *model.User
		wantErr   bool
	}{
		{
			name: "NORMAL: ユーザーを更新できる",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "user",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
				FakeUpdate: func(ctx context.Context, user *model.User) error {
					return nil
				},
			},
			want: &model.User{
				ID:            "id",
				Name:          "user",
				Email:         "email",
				EncryptedPass: "password",
			},
			wantErr: false,
		},
		{
			name: "ERROR: ユーザーのバリデーションに引っかかる場合",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
				FakeUpdate: func(ctx context.Context, user *model.User) error {
					return nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: FindByIDがErrorを返す",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
				FakeUpdate: func(ctx context.Context, user *model.User) error {
					return nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: updateでErrorを返す",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "user",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "user",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
				FakeUpdate: func(ctx context.Context, user *model.User) error {
					return fmt.Errorf("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.UserRepoMock{
				FakeFindByID: tt.repoMocks.FakeFindByID,
				FakeUpdate:   tt.repoMocks.FakeUpdate,
			}
			service := testutils.UserServiceMock{}
			u := &userUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			got, err := u.UpdateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCase.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCase_DeleteUser(t *testing.T) {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	type repoMocks struct {
		FakeFindByID func(id string) (*model.User, error)
		FakeDelete   func(ctx context.Context, user *model.User) error
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name      string
		args      args
		repoMocks repoMocks
		wantErr   bool
	}{
		{
			name: "NORMAL: 正常に削除できる",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "user",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
				FakeDelete: func(ctx context.Context, user *model.User) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "ERROR: FindByUserがErrorを返す",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return nil, fmt.Errorf("error")
				},
				FakeDelete: func(ctx context.Context, user *model.User) error {
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "ERROR: DeleteがErrorを返す",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			repoMocks: repoMocks{
				FakeFindByID: func(id string) (*model.User, error) {
					return &model.User{
						ID:            "id",
						Name:          "user",
						Email:         "email",
						EncryptedPass: "password",
					}, nil
				},
				FakeDelete: func(ctx context.Context, user *model.User) error {
					return fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := testutils.UserRepoMock{
				FakeFindByID: tt.repoMocks.FakeFindByID,
				FakeDelete:   tt.repoMocks.FakeDelete,
			}
			service := testutils.UserServiceMock{}
			u := &userUseCase{
				repo:        repo,
				transaction: transaction,
				service:     service,
			}
			if err := u.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test With DB
func SetUserUseCase() UserUseCase {
	db := testutils.GetTestDB().DB
	transaction := datastore.NewTransaction(db)
	repository := datastore.NewUserRepository(db)
	service := service.NewUserService(repository)
	return NewUserUseCase(repository, transaction, service)
}

func Test_userUseCase_CreateUser_DB(t *testing.T) {
	u := SetUserUseCase()
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		saved   bool
		wantErr bool
	}{
		{
			name: "NORMAL: ユーザーを登録できる",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "name",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			want: &model.User{
				ID:    "id",
				Name:  "name",
				Email: "email", // passwordは返さないようにしている
			},
			saved:   true,
			wantErr: false,
		},
		{
			name: "ERROR: IDが重複した場合は登録できない",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id1",
					Name:          "name",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			want:    nil,
			saved:   false,
			wantErr: true,
		},
		{
			name: "ERROR: Userのバリデーションに引っかかる時も登録できない",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			want:    nil,
			saved:   false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ClearUserData()
			testutils.InsertUserData()
			bCount := testutils.CountUserData()
			got, err := u.CreateUser(tt.args.ctx, tt.args.user)
			aCount := testutils.CountUserData()
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (aCount-bCount == 1) != tt.saved {
				t.Errorf("userUseCase.CreateUser() before Count = %v, after Count = %v, saved: %v", bCount, aCount, tt.saved)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCase.CreateUser() = %v, want %v", got, tt.want)
			}
			testutils.ClearUserData()
		})
	}
}

func Test_userUseCase_VerifyByNameOrEmailPass_DB(t *testing.T) {
	u := SetUserUseCase()
	type args struct {
		ctx         context.Context
		nameOrEmail string
		password    string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "NORMAL: NameOrEmailがDBに存在して, パスワードが正解である(nameで検索)",
			args: args{
				ctx:         context.Background(),
				nameOrEmail: "name1",
				password:    "password1",
			},
			want: &model.User{
				ID:    "id1",
				Name:  "name1",
				Email: "email1",
			},
			wantErr: false,
		},
		{
			name: "NORMAL: NameOrEmailがDBに存在して, パスワードが正解である(emailで検索)",
			args: args{
				ctx:         context.Background(),
				nameOrEmail: "email1",
				password:    "password1",
			},
			want: &model.User{
				ID:    "id1",
				Name:  "name1",
				Email: "email1",
			},
			wantErr: false,
		},
		{
			name: "ERROR: NameOrEmailがDBに存在しない場合",
			args: args{
				ctx:         context.Background(),
				nameOrEmail: "email",
				password:    "password",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: NameOrEmailがDBに存在するが, パスワードが異なる場合",
			args: args{
				ctx:         context.Background(),
				nameOrEmail: "email1",
				password:    "password",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ClearUserData()
			testutils.InsertUserData()
			got, err := u.VerifyByNameOrEmailPass(tt.args.ctx, tt.args.nameOrEmail, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.VerifyByNameOrEmailPass() error = %v, wantErr %v", err, tt.wantErr)
				testutils.ClearUserData()
				return
			}
			if tt.wantErr {
				return
			}
			if (got.ID != tt.want.ID) || (got.Name != tt.want.Name) || (got.Email != tt.want.Email) {
				t.Errorf("userUseCase.VerifyByNameOrEmailPass() = %v, want %v", got, tt.want)
			}
			testutils.ClearUserData()
		})
	}
}

func Test_userUseCase_UpdateUser_DB(t *testing.T) {
	u := SetUserUseCase()
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "NORMAL: ユーザーを更新できる, IDは更新したいもののIDである必要あり",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id1",
					Name:          "name",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			want: &model.User{
				ID:    "id1",
				Name:  "name",
				Email: "email", // passwordは返さないようにしている
			},
			wantErr: false,
		},
		{
			name: "ERROR: IDがDBに存在しない場合は, 更新できない",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "name",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			want:    nil,
			wantErr: false, // FindByIDでは, 存在しないときは, error = nilを返す
		},
		{
			name: "ERROR: Userのバリデーションに引っかかると更新できない",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:            "id",
					Name:          "",
					Email:         "email",
					EncryptedPass: "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ClearUserData()
			testutils.InsertUserData()
			bCount := testutils.CountUserData()
			got, err := u.UpdateUser(tt.args.ctx, tt.args.user)
			aCount := testutils.CountUserData()
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if aCount != bCount {
				t.Errorf("userUseCase.UpdateUser() before Count = %v, after Count = %v", bCount, aCount)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUseCase.UpdateUser() = %v, want %v", got, tt.want)
			}
			testutils.ClearUserData()
		})
	}
}

func Test_userUseCase_DeleteUser_DB(t *testing.T) {
	u := SetUserUseCase()
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		deleted bool
		wantErr bool
	}{
		{
			name: "NORMAL: ユーザーを削除できる",
			args: args{
				ctx: context.Background(),
				id:  "id1",
			},
			deleted: true,
			wantErr: false,
		},
		{
			name: "ERROR: IDがDBに存在しない場合は, 削除できない",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			deleted: false,
			wantErr: false, // FindByIDでは, 存在しないときは, error = nilを返す
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ClearUserData()
			testutils.InsertUserData()
			bCount := testutils.CountUserData()
			err := u.DeleteUser(tt.args.ctx, tt.args.id)
			aCount := testutils.CountUserData()
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (bCount-aCount == 1) != tt.deleted {
				t.Errorf("userUseCase.DeleteUser() before Count = %v, after Count = %v", bCount, aCount)
				return
			}
			testutils.ClearUserData()
		})
	}
}
