package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 定义接口
type IVoteService interface {
	AddVote(c *gin.Context, dto entity.VoteDto)
}

type VoteServiceImpl struct{}

// 投票
func (p *VoteServiceImpl) AddVote(c *gin.Context, dto entity.VoteDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	user, _ := dao.GetUserByUserId(dto.UserId)
	if user.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "投票用户不存在")
		return
	}
	player, _ := dao.GetPlayerDetail(dto.PlayerId)
	if player.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "参赛选手用户不存在")
		return
	}

	vote, _ := dao.GetVoteInfo(dto)

	if vote.ID != 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "已经投过票了")
		return
	}

	err = dao.AddVote(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "投票失败")
	}
	_, err = dao.UpdatePlayer(dto.PlayerId)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "更新选手失败")
	}
	redisKey := "ranking:" + string(rune(player.Aid))
	store.Update(redisKey, int(player.ID))
	result.Success(c, "投票成功")
}

var voteService = VoteServiceImpl{}

func VoteService() IVoteService {
	return &voteService
}
