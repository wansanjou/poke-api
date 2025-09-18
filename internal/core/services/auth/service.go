package authsvc

import (
	"context"
	"errors"
	"time"

	"github.com/wansanjou/poke-api/internal/core/domains"
	"github.com/wansanjou/poke-api/internal/core/ports"
	"github.com/wansanjou/poke-api/utils"
)

type authService struct {
	userrepo ports.UserRepository
}

func NewAuthService(userrepo ports.UserRepository) ports.AuthService {
	return &authService{
		userrepo,
	}
}

func (s *authService) Register(ctx context.Context, data domains.User) (*domains.User, error) {
	if data.Username == "" || data.Password == "" {
		return nil, errors.New("username and password are required")
	}

	user, err := s.userrepo.FindByUsername(ctx, data.Username)
	if user != nil {
		return nil, errors.New("username already exists")
	}
	if err != nil {
		return nil, err
	}

	hash, err := utils.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	data.Password = hash

	return s.userrepo.Create(ctx, data)
}

func (s *authService) Login(ctx context.Context, in domains.LoginRequest) (*domains.LoginResponse, error) {
	if in.Username == "" || in.Password == "" {
		return nil, errors.New("username and password are required")
	}

	user, err := s.userrepo.FindByUsername(ctx, in.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := utils.VerifyPassword(user.Password, in.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.CreateToken(user.Username, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &domains.LoginResponse{Token: token}, nil
}
