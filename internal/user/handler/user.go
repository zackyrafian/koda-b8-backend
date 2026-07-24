package handler

import (
	"belimudah/internal/user/dto"
	"belimudah/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct { 
  service *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler { 
  return &UserHandler{service: service}
}

func (h *UserHandler) Create (c *gin.Context) { 
  var req dto.CreateUserRequest

  if err := c.ShouldBindJSON(&req); err != nil { 
    c.JSON(http.StatusBadRequest, err.Error())
    return 
  }

  id, err := h.service.Create(c.Request.Context(), req) 
  if err != nil { 
    c.JSON(http.StatusInternalServerError, err.Error())
  }
  c.JSON(http.StatusCreated, gin.H{ 
    "id": id,
  })
}