package domain

import (
	"belimudah/internal/product/dto"
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id int64) (*dto.DetailResponse, error)
	FindAll(ctx context.Context) ([]dto.DetailResponse, error)
	Create(ctx context.Context, req dto.CreateRequest) (int64, error)
	Update(ctx context.Context, id int64, req dto.UpdateRequest) (int64, error)
	AddImages(ctx context.Context, productID int64, urls []string) error
	RemoveImages(ctx context.Context, productID int64, imageIDs []int64) error
	GetReviews(ctx context.Context, productID int64) ([]dto.ReviewResponse, error)
	CreateReview(ctx context.Context, productID int64, userID int64, req dto.CreateReviewRequest) error
}