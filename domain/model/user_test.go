package model

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id       string
		name     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "NORMAL: 正常にUserを作成できる",
			args: args{
				id:       "id",
				name:     "name",
				email:    "email",
				password: "password",
			},
			want: &User{
				ID:            "id",
				Name:          "name",
				Email:         "email",
				EncryptedPass: "password",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.id, tt.args.name, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ValidateUser(t *testing.T) {
	type fields struct {
		ID            string
		Name          string
		Email         string
		EncryptedPass string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "NORMAL: 正常な時はnilを返す",
			fields: fields{
				ID:            "id",
				Name:          "name",
				Email:         "email",
				EncryptedPass: "password",
			},
			wantErr: false,
		},
		{
			name: "ERROR: idが空文字の場合",
			fields: fields{
				Name:          "name",
				Email:         "email",
				EncryptedPass: "password",
			},
			wantErr: true,
		},
		{
			name: "ERROR: nameが空文字の場合",
			fields: fields{
				ID:            "id",
				Email:         "email",
				EncryptedPass: "password",
			},
			wantErr: true,
		},
		{
			name: "ERROR: emailが空文字の場合",
			fields: fields{
				ID:            "id",
				Name:          "name",
				EncryptedPass: "password",
			},
			wantErr: true,
		},
		{
			name: "ERROR: passwordが空文字の場合",
			fields: fields{
				ID:    "id",
				Name:  "name",
				Email: "email",
			},
			wantErr: true,
		},
		{
			name: "ERROR: passwordが6文字未満の場合",
			fields: fields{
				ID:            "id",
				Name:          "name",
				Email:         "email",
				EncryptedPass: "12345",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:            tt.fields.ID,
				Name:          tt.fields.Name,
				Email:         tt.fields.Email,
				EncryptedPass: tt.fields.EncryptedPass,
			}
			if err := u.ValidateUser(); (err != nil) != tt.wantErr {
				t.Errorf("User.ValidateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_MatchPassword(t *testing.T) {
	type fields struct {
		ID    string
		Name  string
		Email string
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "NORMAL: パスワードと暗号化されたパスワードが一致するかを",
			fields: fields{
				ID:    "id",
				Name:  "name",
				Email: "email",
			},
			args: args{
				password: "password",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := EncryptPassword(tt.args.password)
			u := &User{
				ID:            tt.fields.ID,
				Name:          tt.fields.Name,
				Email:         tt.fields.Email,
				EncryptedPass: p,
			}
			if got := u.MatchPassword(tt.args.password); got != tt.want {
				t.Errorf("User.MatchPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "NORMAL: パスワードが暗号化されている",
			args: args{
				password: "password",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.args.password {
				t.Errorf("EncryptPassword() = %v is nil", got)
			}
		})
	}
}

func TestConfirmPassword(t *testing.T) {
	type args struct {
		pass        string
		confirmPass string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "NORMAL: パスワードと確認用パスワードが同じだとOK",
			args: args{
				pass:        "password",
				confirmPass: "password",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ERROR: パスワードと確認用パスワードが違う場合",
			args: args{
				pass:        "pass",
				confirmPass: "password",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "ERROR: パスワードが6文字未満の場合",
			args: args{
				pass:        "pass",
				confirmPass: "pass",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "ERROR: パスワードと確認用パスワードが空文字の場合",
			args: args{
				pass:        "",
				confirmPass: "",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfirmPassword(tt.args.pass, tt.args.confirmPass)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfirmPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConfirmPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
