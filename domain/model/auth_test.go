package model

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestNewJWT(t *testing.T) {
	type args struct {
		userID   string
		userName string
	}
	now, _ := time.Parse("2006-01-02 15:04:05 MST", "2020-1-1 00:00:00 JST")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "NORMAL: userID, userName, nowでJWTを作成する",
			args: args{
				userID:   "id",
				userName: "name",
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOi02MjEzNTUxMDQwMCwiaWF0IjotNjIxMzU1OTY4MDAsImlzcyI6IlRhdHN1ZW1vbi10ZXN0IiwibmFtZSI6Im5hbWUiLCJzdWIiOiJpZCJ9.qzyoCwFWeqtJZ1wxGckTZruDTHG9lwcMn8TnPZEH0xQ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJWT(tt.args.userID, tt.args.userName, now)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJWT(t *testing.T) {
	type args struct {
		userID   string
		userName string
		now      time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Auth
		wantErr bool
	}{
		{
			name: "NORMAL: JWTから認証情報を取り出す",
			args: args{
				userID:   "id",
				userName: "name",
				now:      time.Now(),
			},
			want: &Auth{
				UserID:   "id",
				UserName: "name",
				Iss:      os.Getenv("JWT_ISS"),
			},
		},
		{
			name: "NORMAL: idとnameが空文字でもOK",
			args: args{
				userID:   "",
				userName: "",
				now:      time.Now(),
			},
			want: &Auth{
				UserID:   "",
				UserName: "",
				Iss:      os.Getenv("JWT_ISS"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			token, err := NewJWT(tt.args.userID, tt.args.userName, now)
			got, err := ParseJWT(token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.Iat = tt.args.now.Unix()
			tt.want.Exp = tt.args.now.Add(time.Hour * 24).Unix()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJWTWithError(t *testing.T) {
	type args struct {
		token *jwt.Token
	}
	now := time.Now()
	tests := []struct {
		name    string
		args    args
		want    *Auth
		wantErr bool
	}{
		{
			name: "ERROR: NoneのJWTを使用してもParseできない",
			args: args{
				token: jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
					"sub":  "id",
					"name": "name",
					"iss":  os.Getenv("JWT_ISS"),
					"iat":  now.Unix(),
					"exp":  now.Add(time.Hour * 24).Unix(),
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: idがないとErrorを返す",
			args: args{
				token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"name": "name",
					"iss":  os.Getenv("JWT_ISS"),
					"iat":  now.Unix(),
					"exp":  now.Add(time.Hour * 24).Unix(),
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: nameがないとErrorを返す",
			args: args{
				token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"sub": "id",
					"iss": os.Getenv("JWT_ISS"),
					"iat": now.Unix(),
					"exp": now.Add(time.Hour * 24).Unix(),
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: issがないとErrorを返す",
			args: args{
				token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"sub":  "id",
					"name": "name",
					"iat":  now.Unix(),
					"exp":  now.Add(time.Hour * 24).Unix(),
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: iatがないとErrorを返す",
			args: args{
				token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"sub":  "id",
					"name": "name",
					"iss":  os.Getenv("JWT_ISS"),
					"exp":  now.Add(time.Hour * 24).Unix(),
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: expがないとErrorを返す",
			args: args{
				token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"sub":  "id",
					"name": "name",
					"iss":  os.Getenv("JWT_ISS"),
					"iat":  now.Unix(),
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ERROR: 有効期限がきれていたら, Errorを返す",
			args: args{
				token: jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"sub":  "id",
					"name": "name",
					"iss":  os.Getenv("JWT_ISS"),
					"iat":  now.Unix(),
					"exp":  now.Add(time.Hour * -24).Unix(),
				}),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signed, _ := tt.args.token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
			got, err := ParseJWT(signed)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetUserIDInContext(t *testing.T) {
	type args struct {
		parents context.Context
		t       string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "NORMAL: contextにuserIDを保存する",
			args: args{
				parents: context.Background(),
				t:       "user-id",
			},
			want: context.WithValue(context.Background(), userIDContextKey, "user-id"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetUserIDInContext(tt.args.parents, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserIDInContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserIDInContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "NORMAL: contextからuserIDを取得する",
			args: args{
				ctx: context.WithValue(context.Background(), userIDContextKey, "user_id"),
			},
			want:    "user_id",
			wantErr: false,
		},
		{
			name: "ERROR: user_idが設定されていないばあい",
			args: args{
				ctx: context.Background(),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserIDInContext(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserIDInContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserIDInContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
