package auth

import (
	"gitlab.com/url-builder/go-admin/src/database/repositories"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (int, error) {
	password := EncodeSha(a.Password)
	return repositories.AuthRepository.CheckAuth(a.Username, password)
}
