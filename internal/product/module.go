package product

import (
	"belimudah/internal/middleware"
	"belimudah/internal/product/handler"
	"belimudah/internal/product/repository"
	"belimudah/internal/product/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ProductRegister(r gin.IRouter, db *pgxpool.Pool) {
	repo := repository.NewProductRepository(db)
	service := usecase.NewProductService(repo)
	h := handler.NewProductHandler(service)

	r.GET("/products", h.FindAll)
	r.GET("/products/:id", h.FindByID)
	r.GET("/products/:id/reviews", h.GetReviews)

	// auth required for posting reviews
	authOnly := r.Group("/products", middleware.AuthMiddleware())
	authOnly.POST("/:id/reviews", h.CreateReview)

	// auth + admin required for write operations
	protected := r.Group("/products", middleware.AuthMiddleware(), middleware.AdminMiddleware())
	protected.POST("", h.Create)
	protected.PATCH("/:id", h.Update)
	protected.POST("/:id/images", h.AddImages)
	protected.DELETE("/:id/images", h.RemoveImages)
}