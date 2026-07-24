package usecase

import (
	"belimudah/internal/user/domain"
	"belimudah/internal/user/dto"
	"context"
)

type UserService struct { 
  repository domain.Repository
}

func NewUserService(repository domain.Repository) *UserService { 
  return &UserService{repository: repository}
}

func (s *UserService) Create(ctx context.Context, req dto.CreateUserRequest) (int64, error) { 
  return s.repository.Create(ctx, req)
}
