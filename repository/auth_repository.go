package repository

import (
	"context"
	"fmt"

	"github.com/haleyrc/idealize/domain"
	"github.com/haleyrc/idealize/internal/pg"
	"github.com/haleyrc/idealize/repository/queries"
)

type AuthRepository struct{}

func (repo *AuthRepository) GetUser(ctx context.Context, tx pg.Tx, u *domain.User, id string) error {
	if err := tx.QueryRow(ctx, queries.GetUserQuery, id).Scan(&u.ID, &u.Username); err != nil {
		return fmt.Errorf("get user failed: %w", err)
	}
	return nil
}
