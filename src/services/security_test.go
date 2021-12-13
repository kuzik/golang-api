package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
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
			want: "03a7c6327685ae2a8c745bd62a02848adb286b6648835f364aea3dcb7940a148",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SecurityService{Secret: tt.args.secret}
			if got := s.EncodeSha(tt.args.password); got != tt.want {
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
		{
			name: "Happy flow",
			args: args{
				username: "user1",
				userId:   1,
				secret:   "secret",
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwidXNlcl9pZCI6MSwiZXhwIjoxODkzNDU2MDAwLCJpc3MiOiJnaW4tYmxvZyJ9.z3LuwJwFvtMrs8sD_xUQd9lDck5KaWUZT0aqV5TM47Q",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SecurityService{Secret: tt.args.secret}
			ttl, _ := time.Parse("2006/01/02", "2030/01/01")
			got, err := s.GenerateToken(tt.args.username, tt.args.userId, ttl)
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
	ttl, _ := time.Parse("2006/01/02", "2030/01/01")
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		{
			name: "Happy flow",
			args: args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwidXNlcl9pZCI6MSwiZXhwIjoxODkzNDU2MDAwLCJpc3MiOiJnaW4tYmxvZyJ9.z3LuwJwFvtMrs8sD_xUQd9lDck5KaWUZT0aqV5TM47Q"},
			want: &Claims{
				Username: "user1",
				UserId:   1,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: ttl.Unix(),
					Issuer:    "gin-blog",
				},
			},
		},
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
