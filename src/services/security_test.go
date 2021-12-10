package services

import (
	"reflect"
	"testing"
)

func TestSecurityService_EncodeSha(t *testing.T) {
	type args struct {
		password string
		secret   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Sha test",
			args: args{"password", "1111"},
			want: "7f777bccedb2356c3ebab81ddf752e65664ee4248e356e25c3310919860eeb02",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SecurityService{}
			if got := s.EncodeSha(tt.args.password, tt.args.secret); got != tt.want {
				t.Errorf("EncodeSha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurityService_GenerateToken(t *testing.T) {
	type args struct {
		username string
		password string
		userId   int
		secret   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SecurityService{}
			got, err := s.GenerateToken(tt.args.username, tt.args.password, tt.args.userId, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurityService_ParseToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SecurityService{}
			got, err := s.ParseToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
