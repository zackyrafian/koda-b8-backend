package dto

type CreateCategoryRequest struct { 
  Name string `json:"name"`
}

type CategoryResponse struct { 
  Id int64 `json:"id"`
  Name string `json:"name"`
}