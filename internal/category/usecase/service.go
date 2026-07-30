package usecase

import (
	"belimudah/internal/category/domain"
	"belimudah/internal/category/dto"
	"context"

)

type Service struct { 
  repository domain.Repository
}

func NewCategoryService(repository domain.Repository) *Service { 
  return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, req dto.CreateCategoryRequest) (int64, error) { 
  return s.repository.Create(ctx, req)
}

func (s *Service) FindAll(ctx context.Context) ([]domain.Category, error) {
  return s.repository.GetAll(ctx)
}

func (s *Service) FindByID(ctx context.Context, id int64) (domain.Category, error) { 
  return s.repository.GetByID(ctx, id)
}