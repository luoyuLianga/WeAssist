package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"

	"github.com/gin-gonic/gin"
)

// 投票
// @Summary 投票接口
// @Tags      Vote
// @Produce json
// @Description 投票接口
// @Param data body entity.VoteDto true "data"
// @Success 200 {object} result.Result
// @router /api/vote/add [post]
func AddVote(c *gin.Context) {
	var dto entity.VoteDto
	_ = c.BindJSON(&dto)
	service.VoteService().AddVote(c, dto)
}
