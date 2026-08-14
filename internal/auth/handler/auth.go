package handler

import (
	"belimudah/internal/auth/dto"
	"belimudah/internal/auth/usecase"
	"fmt"
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
  name, err := h.service.Register(c.Request.Context(), req)
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "success": false,
      "error": err.Error(),
    })
    return
  }
  c.JSON(http.StatusCreated, gin.H{ 
    "success": true, 
    "message": fmt.Sprintf("Successfully created account %s",  name),
  })
}

func (h *AuthHandler) Login(c *gin.Context) { 
  var req dto.LoginRequest

  if err := c.ShouldBindJSON(&req); err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "success": false,
      "message": err.Error(),
    })
    return
  }

  fmt.Print(req)
  token, err := h.service.Login(c.Request.Context(), req) 
  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "success": false,
      "message": err.Error(),
    })
    return 
  }
  c.JSON(http.StatusAccepted, gin.H{ 
    "success": true,
    "token": token,
  })
}

func (h *AuthHandler) ForgetPassword(c *gin.Context) { 
  var email string

  if err := c.ShouldBindJSON(&email); err != nil { 
    c.JSON(http.StatusBadRequest, gin.H{ 
      "success": false, 
      "message": "Failed format email", 
    })
  }

  email, err := h.service.ForgetPassword(c.Request.Context(), email)

  if err != nil { 
    c.JSON(http.StatusInternalServerError, gin.H{ 
      "success": false, 
      "message": "",
    })
  }
  c.JSON(http.StatusOK, gin.H{ 
    "success": true,
    "data" : email, 
  })
}