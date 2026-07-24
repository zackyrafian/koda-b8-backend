package dto

type CreateRequest struct {
	Name        string   `json:"name" binding:"required"`
	BrandID     int64    `json:"brand_id" binding:"required"`
	CategoryID  int64    `json:"category_id" binding:"required"`
	Price       float64  `json:"price" binding:"required"`
	Discount    float64  `json:"discount"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`

	Variants []string `json:"variants"`
	Images   []string `json:"images"`
}