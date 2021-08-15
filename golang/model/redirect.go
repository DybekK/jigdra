package model

import (
	"context"
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
	return r.repo.SecureRedirect(ctx, id)
}

func (r *redirectService) VerifyRedirect(ctx context.Context, hex string) (string, error) {
	return r.repo.VerifyRedirect(ctx, hex)
}
