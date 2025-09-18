package config

import (
	"errors"
	"log"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server Server
	Mongo  Mongo
	Jwt    JWT
}

type Server struct {
	Port int `envconfig:"PORT" default:"8080"`
}

type Mongo struct {
	URI      string `envconfig:"MONGO_URI" default:"mongodb://localhost:27017"`
	Database string `envconfig:"DB_NAME" default:"poke-api"`
}

type JWT struct {
	SecretKey string `envconfig:"JWT_SECRET_KEY" required:"true"`
}

var cfg Config

func Init() {
	runtime.GOMAXPROCS(1)

	_ = godotenv.Load()

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error: %s", err.Error())
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("config validation error: %s", err)
	}

	log.Printf("Configuration loaded successfully")
}

func Get() Config {
	return cfg
}

func (c *Config) Validate() error {
	if c.Jwt.SecretKey == "" {
		return errors.New("JWT secret key is required")
	}

	if len(c.Jwt.SecretKey) < 32 {
		return errors.New("JWT secret key should be at least 32 characters long")
	}

	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return errors.New("server port must be between 1-65535")
	}

	return nil
}
