package handler

import (
	"belimudah/internal/brand/dto"
	"belimudah/internal/brand/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct { 
  service *usecase.Service
}

func NewBrandHander(service *usecase.Service) *Handler { 
  return &Handler{service: service}
}

func (h *Handler) Create (c *gin.Context){ 
  var req dto.BrandCreateRequest
  if err := c.ShouldBindJSON(&req); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
          "error": err.Error(),
      })
      return
  }

  id, err := h.service.Create(c.Request.Context(), req)
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "error": err.Error(),
    })
    return
  }
  c.JSON(http.StatusCreated, gin.H{ 
    "id": id,
  })
}