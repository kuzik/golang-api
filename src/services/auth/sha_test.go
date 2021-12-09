package auth_test

import (
	"testing"

	"gitlab.com/url-builder/go-admin/src/services/auth"
)

func TestEncodeSha(t *testing.T) {
	t.Run("Sha testing", func(t *testing.T) {
		expected := "7f777bccedb2356c3ebab81ddf752e65664ee4248e356e25c3310919860eeb02"
		if got := auth.EncodeSha("password"); got != expected {
			t.Errorf("EncodeSha() = %v, want %v", got, expected)
		}
	})
}
