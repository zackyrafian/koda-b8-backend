package usecase

import (
	"belimudah/internal/brand/domain"
	"belimudah/internal/brand/dto"
	"context"
)

type Service struct { 
  repo domain.Repository
}

func NewServiceBrand(repo domain.Repository) *Service { 
  return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req dto.BrandCreateRequest) (int64, error) { 
  return s.repo.Create(ctx, req)
}