package service

import (
	"context"

	"github.com/haleyrc/idealize/domain"
	"github.com/haleyrc/idealize/internal/logger"
)

var _ domain.AuthService = &authLogger{}

type authLogger struct {
	Logger  logger.Logger
	Service domain.AuthService
}

func (l *authLogger) GetUser(ctx context.Context, id string) (*domain.User, error) {
	l.Logger.Debug(ctx, "getting user", logger.Fields{"ID": id})

	user, err := l.Service.GetUser(ctx, id)
	if err != nil {
		l.Logger.Error(ctx, "get user failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "got user", logger.Fields{"User": user})

	return user, nil
}
