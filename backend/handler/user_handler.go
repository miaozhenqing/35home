package handler

import (
	"net/http"

	"35home/common"
	"35home/dto"
	"35home/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponseWithBody(common.ErrInvalidRequest, err.Error()))
		return
	}

	response, err := h.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(common.ErrInternalServer))
		return
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponse(response))
}

func (h *userHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(common.ErrInvalidRequest))
		return
	}

	response, err := h.userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(common.ErrInvalidEmailOrPassword))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}
