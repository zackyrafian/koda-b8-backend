package user

import (
	"belimudah/internal/user/handler"
	"belimudah/internal/user/repository"
	"belimudah/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UserRegister(r gin.IRouter, db *pgxpool.Pool) { 
  repo := repository.NewUserRepository(db)
  service := usecase.NewUserService(repo) 
  handler := handler.NewUserHandler(service) 

  r.POST("/user", handler.Create)
}