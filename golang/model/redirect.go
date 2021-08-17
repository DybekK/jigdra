package model

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"golang/model/dto"
	"golang/model/repository"
)

type redirectService struct {
	repo repository.RedirectRepository
}

type RedirectService interface {
	SecureRedirect(context.Context, string) (string, error)
	VerifyRedirect(context.Context, string) (string, error)
}

func NewRedirectService(repo repository.RedirectRepository) RedirectService {
	return &redirectService{repo: repo}
}

func (r *redirectService) SecureRedirect(ctx context.Context, id string) (string, error) {
	var sec dto.Security
	sec.Id = id
	randHex, _ := randomHex(20)
	sec.Hex = randHex
	return r.repo.SecureRedirect(ctx, sec)
}

func (r *redirectService) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	return r.repo.VerifyRedirect(ctx, hex)
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
