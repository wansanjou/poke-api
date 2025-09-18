package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wansanjou/poke-api/internal/core/domains"
	"github.com/wansanjou/poke-api/internal/core/ports"
)

type pokeClient struct {
	http *http.Client
	base string
}

func NewPokemonClient() ports.PokemonClient {
	return &pokeClient{
		http: &http.Client{Timeout: 10 * time.Second},
		base: "https://pokeapi.co/api/v2",
	}
}

func (c *pokeClient) GetPokemon(name string) (*domains.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", c.base, name)
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pokeapi error: %s", resp.Status)
	}

	var raw domains.PokemonResponse
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	pokemon := &domains.Pokemon{
		Name:   raw.Name,
		Weight: raw.Weight,
		Height: raw.Height,
	}

	for _, t := range raw.Types {
		pokemon.Types = append(pokemon.Types, t.Type.Name)
	}
	for _, a := range raw.Abilities {
		pokemon.Abilities = append(pokemon.Abilities, a.Ability.Name)
	}

	return pokemon, nil
}
