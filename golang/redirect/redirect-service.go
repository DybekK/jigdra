package redirect

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type RedirectService struct {
	repo RedirectRepository
}

//factory
func NewRedirectService(repo RedirectRepository) RedirectService {
	return RedirectService{repo: repo}
}

//methods
func (rs *RedirectService) SecureRedirect(ctx context.Context, id string) (string, error) {
	var sec Security
	sec.Id = id
	randHex, _ := randomHex(20)
	sec.Hex = randHex
	return rs.repo.SecureRedirect(ctx, sec)
}

func (rs *RedirectService) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	return rs.repo.VerifyRedirect(ctx, hex)
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
