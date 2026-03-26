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
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:     "INVALID_REQUEST",
			Msg:      "Invalid request parameters",
			RespBody: err.Error(),
		})
		return
	}

	response, err := h.userService.Register(&req)
	if err != nil {
		if err.Error() == common.ErrorEmailAlreadyRegistered {
			c.JSON(http.StatusConflict, dto.Response{
				Code: common.ErrorEmailAlreadyRegistered,
				Msg:  common.ErrorMessages[common.ErrorEmailAlreadyRegistered],
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code: err.Error(),
			Msg:  common.ErrorMessages[err.Error()],
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Code:     "SUCCESS",
		Msg:      "User registered successfully",
		RespBody: response,
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: "INVALID_REQUEST",
			Msg:  "Invalid request parameters",
		})
		return
	}

	response, err := h.userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Code: err.Error(),
			Msg:  common.ErrorMessages[err.Error()],
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:     "SUCCESS",
		Msg:      "Login successful",
		RespBody: response,
	})
}
