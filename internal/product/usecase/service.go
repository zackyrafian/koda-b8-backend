package usecase

import (
	"belimudah/internal/product/domain"
	"belimudah/internal/product/dto"
	"context"
	"errors"
)

type Service struct {
	repository domain.Repository
}

func NewProductService(repository domain.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) FindByID(ctx context.Context, id int64) (*dto.DetailResponse, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *Service) FindAll(ctx context.Context) ([]dto.DetailResponse, error) {
	return s.repository.FindAll(ctx)
}

func (s *Service) Create(ctx context.Context, req dto.CreateRequest) (int64, error) {
	if req.Price < 0 {
		return 0, errors.New("price must be positive")
	}
	if req.Stock < 0 {
		return 0, errors.New("stock must be positive")
	}
	return s.repository.Create(ctx, req)
}

func (s *Service) Update(ctx context.Context, id int64, req dto.UpdateRequest) (int64, error) {
	return s.repository.Update(ctx, id, req)
}

func (s *Service) AddImages(ctx context.Context, productID int64, urls []string) error {
	return s.repository.AddImages(ctx, productID, urls)
}

func (s *Service) RemoveImages(ctx context.Context, productID int64, imageIDs []int64) error {
	return s.repository.RemoveImages(ctx, productID, imageIDs)
}

func (s *Service) GetReviews(ctx context.Context, productID int64) ([]dto.ReviewResponse, error) {
	return s.repository.GetReviews(ctx, productID)
}

func (s *Service) CreateReview(ctx context.Context, productID int64, userID int64, req dto.CreateReviewRequest) error {
	return s.repository.CreateReview(ctx, productID, userID, req)
}
