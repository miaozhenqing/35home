package handler

import (
	"net/http"

	"35home/common"
	"35home/dto"
	"35home/service"

	"github.com/gin-gonic/gin"
)

type VoteHandler interface {
	SubmitVote(c *gin.Context)
	GetVoteStats(c *gin.Context)
}

type voteHandler struct {
	voteService service.VoteService
}

func NewVoteHandler(voteService service.VoteService) VoteHandler {
	return &voteHandler{voteService: voteService}
}

func (h *voteHandler) SubmitVote(c *gin.Context) {
	// 从上下文中获取用户ID（由JWT中间件设置）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(common.ErrInvalidRequest))
		return
	}

	var req dto.VoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(common.ErrInvalidRequest))
		return
	}

	response, err := h.voteService.SubmitVote(userID.(uint64), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(common.ErrInternalServer))
		return
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponse(response))
}

func (h *voteHandler) GetVoteStats(c *gin.Context) {
	response, err := h.voteService.GetVoteStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(common.ErrInternalServer))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(response))
}
