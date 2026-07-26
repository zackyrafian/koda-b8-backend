package handler

import (
	"belimudah/internal/auth/dto"
	"belimudah/internal/auth/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct { 
  service *usecase.AuthService
}

func NewAuthHandler (service *usecase.AuthService) *AuthHandler{ 
  return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) { 
  var req dto.RegisterRequest
  if err := c.ShouldBindJSON(&req); err != nil { 
    c.JSON(http.StatusBadRequest, err.Error())
    return
  }

  id, err := h.service.Register(c.Request.Context(), req)
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