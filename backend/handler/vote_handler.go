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
	// 模拟用户ID，实际项目中应该从JWT token中获取
	userID := uint(1)

	var req dto.VoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: "INVALID_REQUEST",
			Msg:  "Invalid request parameters",
		})
		return
	}

	response, err := h.voteService.SubmitVote(userID, &req)
	if err != nil {
		if err.Error() == common.ErrorUserAlreadyVoted {
			c.JSON(http.StatusConflict, dto.Response{
				Code: common.ErrorUserAlreadyVoted,
				Msg:  common.ErrorMessages[common.ErrorUserAlreadyVoted],
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
		Msg:      "Vote submitted successfully",
		RespBody: response,
	})
}

func (h *voteHandler) GetVoteStats(c *gin.Context) {
	response, err := h.voteService.GetVoteStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code: err.Error(),
			Msg:  common.ErrorMessages[err.Error()],
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:     "SUCCESS",
		Msg:      "Get vote stats successful",
		RespBody: response,
	})
}
