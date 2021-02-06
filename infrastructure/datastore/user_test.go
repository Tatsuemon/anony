package datastore

import (
	"context"
	"reflect"
	"testing"

	"github.com/Tatsuemon/anony/domain/model"
	"github.com/Tatsuemon/anony/domain/repository"
	"github.com/jmoiron/sqlx"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		conn *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want repository.UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindAll(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindByID(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindByName(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.FindByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindByEmail(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.FindByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindByNameOrEmail(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		nameOrEmail string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.FindByNameOrEmail(tt.args.nameOrEmail)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindByNameOrEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindByNameOrEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindDuplicatedUsers(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.FindDuplicatedUsers(tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindDuplicatedUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindDuplicatedUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Store(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.Store(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Update(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			got, err := r.Update(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Delete(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				conn: tt.fields.conn,
			}
			if err := r.Delete(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
