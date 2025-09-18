package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wansanjou/poke-api/internal/core/ports"
)

type pokemonHandler struct {
	pksvc ports.PokemonService
}

func NewPokemonHandler(pksvc ports.PokemonService) *pokemonHandler {
	return &pokemonHandler{
		pksvc,
	}
}

func (h *pokemonHandler) GetPokemon(c *fiber.Ctx) error {
	name := c.Params("name")
	p, err := h.pksvc.GetPokemon(name)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(p)
}

func (h *pokemonHandler) GetAbilities(c *fiber.Ctx) error {
	name := c.Params("name")
	abilities, err := h.pksvc.GetAbilities(name)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"name": name, "abilities": abilities})
}

func (h *pokemonHandler) GetRandom(c *fiber.Ctx) error {
	pokemon, err := h.pksvc.GetRandomPokemon()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(pokemon)
}
