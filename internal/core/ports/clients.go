package ports

import "github.com/wansanjou/poke-api/internal/core/domains"

type PokemonClient interface {
	GetPokemon(name string) (*domains.Pokemon, error)
}
