package dto

type CreateUserRequest struct { 
  FullName string `json:"fullname"`
  Email string `json:"email"`
  Password string `json:"password"`
}
