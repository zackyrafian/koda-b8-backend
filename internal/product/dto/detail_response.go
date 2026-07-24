package dto

type DetailResponse struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	BrandID     int64    `json:"brand_id"`
	Brand       string   `json:"brand"`
	CategoryID  int64    `json:"category_id"`
	Category    string   `json:"category"`
	Price       int64    `json:"price"`
	Discount    int      `json:"discount"`
	Rating      float64  `json:"rating"`
	Stock       int      `json:"stock"`
	SoldOut     int64    `json:"sold_out"`
	Description string   `json:"description"`
	Variants []string `json:"variants"`
	Images   []string `json:"images"`
}