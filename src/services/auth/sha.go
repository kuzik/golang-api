package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"gitlab.com/url-builder/go-admin/src/config"
)

func EncodeSha(value string) string {
	m := hmac.New(sha256.New, []byte(config.App.Secret))
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
