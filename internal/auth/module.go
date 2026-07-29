package auth

import (
	"belimudah/internal/auth/handler"
	"belimudah/internal/auth/repository"
	"belimudah/internal/auth/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRegiter(r gin.IRouter, db *pgxpool.Pool) { 
  repo := repository.NewAuthRepository(db)
  service := usecase.NewAuthService(repo)
  handler := handler.NewAuthHandler(service)

  auth := r.Group("/auth")
  auth.POST("/register", handler.Register)
  auth.POST("/login", handler.Login)
}