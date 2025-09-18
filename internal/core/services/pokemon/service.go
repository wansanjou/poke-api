package pokemonsvc

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/wansanjou/poke-api/internal/core/domains"
	"github.com/wansanjou/poke-api/internal/core/ports"
)

type pokemonService struct {
	client ports.PokemonClient
	cache  *cache.Cache
}

func NewPokemonService(client ports.PokemonClient) ports.PokemonService {
	c := cache.New(10*time.Minute, 15*time.Minute)
	return &pokemonService{
		client: client,
		cache:  c,
	}
}

func (s *pokemonService) GetPokemon(name string) (*domains.Pokemon, error) {
	if cache, found := s.cache.Get(name); found {
		return cache.(*domains.Pokemon), nil
	}

	pokemon, err := s.client.GetPokemon(name)
	if err != nil {
		return nil, err
	}

	s.cache.Set(name, pokemon, cache.DefaultExpiration)
	return pokemon, nil
}

func (s *pokemonService) GetAbilities(name string) ([]string, error) {
	cacheKey := name + "_abilities"
	if x, found := s.cache.Get(cacheKey); found {
		return x.([]string), nil
	}

	p, err := s.GetPokemon(name)
	if err != nil {
		return nil, err
	}

	s.cache.Set(cacheKey, p.Abilities, cache.DefaultExpiration)
	return p.Abilities, nil
}

func (s *pokemonService) GetRandomPokemon() (*domains.Pokemon, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=2000")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pokeapi error: %s", resp.Status)
	}

	var list struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}

	if len(list.Results) == 0 {
		return nil, fmt.Errorf("no pokemon found")
	}

	idx := rand.Intn(len(list.Results))
	name := list.Results[idx].Name

	return s.client.GetPokemon(name)
}
