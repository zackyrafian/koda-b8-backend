package domain

import "time"

type User struct { 
  ID int64 
  Fullname string
  Email string 
  Password string 
  CreatedAt time.Time
}

type UserResponse struct { 
  ID int64  `json:"id"`
  Fullname string `json:"fullname"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"created_at"`
}