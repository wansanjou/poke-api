package ports

import (
	"context"

	"github.com/wansanjou/poke-api/internal/core/domains"
)

type AuthService interface {
	Register(ctx context.Context, data domains.User) (*domains.User, error)
	Login(ctx context.Context, in domains.LoginRequest) (*domains.LoginResponse, error)
}

type PokemonService interface {
	GetPokemon(name string) (*domains.Pokemon, error)
	GetAbilities(name string) ([]string, error)
	GetRandomPokemon() (*domains.Pokemon, error)
}
