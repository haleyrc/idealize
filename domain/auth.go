package domain

import "context"

type AuthService interface {
	GetUser(ctx context.Context, id string) (*User, error)
}
