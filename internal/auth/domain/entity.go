package domain

import (
	"belimudah/internal/auth/dto"
	"context"
)



type Repository interface { 
  Register(ctx context.Context, req dto.RegisterRequest) (int64, error)
  Login(ctx context.Context, req dto.LoginRequest) (error)
  ForgetPassword(ctx context.Context, req dto.ForgetPassword) (string, error)
}