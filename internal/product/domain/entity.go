package domain

import (
	Brand "belimudah/internal/brand/domain"
	Category "belimudah/internal/category/domain"
	"time"
)

type Product struct {
	ID int64
	Name string
	BrandID int64
	CategoryID int64
	Price int64
	Discount int
	Rating float64
	Stock int
	SoldOut int64
	Description string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type ProductDetail struct { 
  Product
  Images ProductImage
  Variants ProductVariant
  Brand Brand.Brand
  Category Category.Category
}