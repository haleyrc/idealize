package service

import (
	"context"
	"fmt"

	"github.com/haleyrc/idealize/domain"
	"github.com/haleyrc/idealize/internal/logger"
	"github.com/haleyrc/idealize/internal/pg"
)

var _ domain.AuthService = &authService{}

type AuthRepository interface {
	GetUser(ctx context.Context, tx pg.Tx, u *domain.User, id string) error
}

func NewAuthService(db Database, logger logger.Logger, repo AuthRepository) domain.AuthService {
	return &authLogger{
		Logger: logger,
		Service: &authService{
			DB:   db,
			Repo: repo,
		},
	}
}

type authService struct {
	DB   Database
	Repo AuthRepository
}

func (svc *authService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		return svc.Repo.GetUser(ctx, tx, &user, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get user fialed: %w", err)
	}

	return &user, nil
}
