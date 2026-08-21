package dto

type UpdateRequest struct {
	Name        string   `json:"name"`
	BrandID     int64    `json:"brand_id"`
	CategoryID  int64    `json:"category_id"`
	Price       float64  `json:"price"`
	Discount    float64  `json:"discount"`
	Stock       int      `json:"stock"`
	Description string   `json:"description"`
	Variants    []string `json:"variants"`
}

type AddImagesRequest struct {
	URLs []string `json:"urls" binding:"required"`
}

type RemoveImagesRequest struct {
	ImageIDs []int64 `json:"image_ids" binding:"required"`
}

type CreateReviewRequest struct {
	VariantID int64   `json:"variant_id" binding:"required"`
	Rating    float64 `json:"rating" binding:"required,min=1,max=5"`
	Comment   string  `json:"comment"`
}

type ReviewResponse struct {
	ID        int64   `json:"id"`
	UserID    int64   `json:"user_id"`
	Fullname  string  `json:"fullname"`
	VariantID int64   `json:"variant_id"`
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
}
