package config

import "testing"

func TestDSN(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "NORMAL: 正常な値",
			want: "test-user:test-pass@tcp(db-test:3306)/anonyURL-test?parseTime=true&collation=utf8mb4_bin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DSN(); got != tt.want {
				t.Errorf("DSN() = %v, want %v", got, tt.want)
			}
		})
	}
}
