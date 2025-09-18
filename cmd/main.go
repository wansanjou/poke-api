package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wansanjou/poke-api/config"
	"github.com/wansanjou/poke-api/infrastructures"
	"github.com/wansanjou/poke-api/internal/clients"
	authsvc "github.com/wansanjou/poke-api/internal/core/services/auth"
	pokemonsvc "github.com/wansanjou/poke-api/internal/core/services/pokemon"
	"github.com/wansanjou/poke-api/internal/handlers"
	"github.com/wansanjou/poke-api/internal/repositories"
	"github.com/wansanjou/poke-api/middleware"
)

func main() {

	config.Init()
	cfg := config.Get()
	db := infrastructures.NewMongoDB()

	userRepo := repositories.NewUserRepository(db, cfg.Mongo.Database)
	authsvc := authsvc.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authsvc)

	pokeCli := clients.NewPokemonClient()
	pokeSvc := pokemonsvc.NewPokemonService(pokeCli)
	pokeHandler := handlers.NewPokemonHandler(pokeSvc)

	app := fiber.New()

	api := app.Group("/api/v1")

	//auth routes
	public := api.Group("/auth")
	public.Post("/register", authHandler.Register)
	public.Post("/login", authHandler.Login)

	//pokemon routes
	protected := api.Group("/pokemon")
	protected.Use(middleware.JWTMiddleware())
	protected.Get("/random", pokeHandler.GetRandom)
	protected.Get("/:name/ability", pokeHandler.GetAbilities)
	protected.Get("/:name", pokeHandler.GetPokemon)

	log.Println("Server started on port", cfg.Server.Port)
	log.Fatal(app.Listen(":" + fmt.Sprintf("%d", cfg.Server.Port)))
}
