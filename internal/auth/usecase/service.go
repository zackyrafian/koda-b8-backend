package usecase

import (
	"belimudah/internal/auth/dto"
	"belimudah/internal/auth/repository"
	"belimudah/internal/libs"
	"context"
	"errors"
	"fmt"
)

type AuthService struct { 
  repository *repository.AuthRepository
}

func NewAuthService(repository *repository.AuthRepository) *AuthService { 
  return &AuthService{repository: repository}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (string, error) { 
  if len(req.Password) < 8 { 
    return "", errors.New("Minimum password length 8")
  }
  hashPassword, err := libs.HashPassword(req.Password)
  if err != nil { 
    return "", err
  }
  user := dto.RegisterRequest{ 
    FullName: req.FullName,
    Email: req.Email, 
    Password: hashPassword,
  }

  return s.repository.Register(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (string, error) { 
  user, err := s.repository.Login(ctx, req)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}
  
	ok, err := libs.Verify(req.Password, user.Password)
	if err != nil {
		return "", err
	}
  
	if !ok {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := libs.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) ForgetPassword(ctx context.Context, email string) (string, error) { 
  // send email
  // 
  return "", nil
}

