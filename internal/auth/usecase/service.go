package usecase

import (
	"belimudah/internal/auth/dto"
	"belimudah/internal/auth/repository"
	userDomain "belimudah/internal/user/domain"
	"context"
	"errors"
)

type AuthService struct { 
  repository *repository.AuthRepository
}

func NewAuthService(repository *repository.AuthRepository) *AuthService { 
  return &AuthService{repository: repository}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (int64, error) { 
  if len(req.Password) < 8 { 
    return 0, errors.New("")
  }
  return s.repository.Register(ctx, req)
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (userDomain.User, error) { 

  return s.repository.Login(ctx, req)
}



