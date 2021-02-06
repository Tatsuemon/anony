package model

import (
	"reflect"
	"testing"
)

func TestNewAnonyURL(t *testing.T) {
	type args struct {
		id       string
		original string
		short    string
		status   int64
	}
	tests := []struct {
		name    string
		args    args
		want    *AnonyURL
		wantErr bool
	}{
		{
			name: "NORMAL: 正常にAnonyURLを作成できる",
			args: args{
				id:       "test-id",
				original: "original-url",
				short:    "short-url",
				status:   1,
			},
			want: &AnonyURL{
				ID:       "test-id",
				Original: "original-url",
				Short:    "short-url",
				Status:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAnonyURL(tt.args.id, tt.args.original, tt.args.short, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAnonyURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnonyURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnonyURL_ValidateAnonyURL(t *testing.T) {
	type fields struct {
		ID       string
		Original string
		Short    string
		Status   int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "NORMAL: 正常な場合は, nilを返す",
			fields: fields{
				ID:       "id",
				Original: "original",
				Short:    "short",
				Status:   1,
			},
			wantErr: false,
		},
		{
			name: "ERROR: IDがない場合",
			fields: fields{
				Original: "original",
				Short:    "short",
				Status:   1,
			},
			wantErr: true,
		},
		{
			name: "ERROR: Originalがない場合",
			fields: fields{
				ID:     "id",
				Short:  "short",
				Status: 1,
			},
			wantErr: true,
		},
		{
			name: "ERROR: Shortがない場合",
			fields: fields{
				ID:       "id",
				Original: "original",
				Status:   1,
			},
			wantErr: true,
		},
		{
			name: "ERROR: Statusがない場合",
			fields: fields{
				ID:       "id",
				Original: "original",
				Short:    "short",
			},
			wantErr: true,
		},
		{
			name: "ERROR: Statusが1未満の場合",
			fields: fields{
				ID:       "id",
				Original: "original",
				Short:    "short",
				Status:   0,
			},
			wantErr: true,
		},
		{
			name: "ERROR: Statusが2より大きい場合",
			fields: fields{
				ID:       "id",
				Original: "original",
				Short:    "short",
				Status:   3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AnonyURL{
				ID:       tt.fields.ID,
				Original: tt.fields.Original,
				Short:    tt.fields.Short,
				Status:   tt.fields.Status,
			}
			if err := a.ValidateAnonyURL(); (err != nil) != tt.wantErr {
				t.Errorf("AnonyURL.ValidateAnonyURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
