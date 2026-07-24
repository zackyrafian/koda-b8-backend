package domain

import (
	"belimudah/internal/user/dto"
	"context"
)

type Repository interface { 
  Create(ctx context.Context, req dto.CreateUserRequest) (int64, error) 
}