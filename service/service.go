package service

import (
	"context"

	"github.com/haleyrc/idealize/internal/pg"
)

type Database interface {
	WithTx(context.Context, func(context.Context, pg.Tx) error) error
}
