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

func (s *Service) FindAll(ctx context.Context) ([]domain.Brand, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) FindByID(ctx context.Context, id int64) (domain.Brand, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, req dto.BrandCreateRequest) (int64, error) {
	return s.repo.Create(ctx, req)
}

func (s *Service) Update(ctx context.Context, id int64, req dto.BrandCreateRequest) (domain.Brand, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *Service) Delete(ctx context.Context, id int64) (string, error) {
	return s.repo.Delete(ctx, id)
}
