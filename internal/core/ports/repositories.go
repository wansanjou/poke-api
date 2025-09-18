package ports

import (
	"context"

	"github.com/wansanjou/poke-api/internal/core/domains"
)

type UserRepository interface {
	Create(ctx context.Context, data domains.User) (*domains.User, error)
	FindByUsername(ctx context.Context, username string) (*domains.User, error)
}
