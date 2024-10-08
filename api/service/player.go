// 用户服务层
package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"WeAssist/common/util"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var store = util.RedisStore{}

// 定义接口
type IPlayerService interface {
	GetPlayerList(c *gin.Context, dto entity.PlayerListDto)
	GetRankList(c *gin.Context, dto entity.PlayerListDto)
	AddPlayer(c *gin.Context, dto entity.AddPlayerDto)
	GetPlayerById(c *gin.Context, id int)
}

type PlayerServiceImpl struct{}

// 获取列表
func (p *PlayerServiceImpl) GetPlayerList(c *gin.Context, dto entity.PlayerListDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	ret, err := dao.GetPlayerList(dto.Aid, "id asc")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "获取选手列表失败")
	}
	result.Success(c, ret)
}

// 获取排行
func (p *PlayerServiceImpl) GetRankList(c *gin.Context, dto entity.PlayerListDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	redisKey := "ranking:" + string(rune(dto.Aid))
	cacheList := store.Get(redisKey)
	fmt.Println(cacheList)
	if len(cacheList) > 0 {
		var players []entity.Player
		// var ids []int
		for _, value := range cacheList {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := dao.GetPlayerDetail(id)
			if rsInfo.ID > 0 {
				players = append(players, rsInfo)
			}
			// ids = append(ids, id)
		}
		// ret, err := dao.GetPlayerByIds(ids)
		if err != nil {
			result.Failed(c, int(result.ApiCode.FAILED), "获取选手列表失败")
		}
		result.Success(c, players)
		return
	}

	ret, err := dao.GetPlayerList(dto.Aid, "score desc")

	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "获取选手列表失败")
	}
	for _, value := range ret {
		store.Set(redisKey, int(value.ID), value.Score)
	}
	result.Success(c, ret)
}

// 新增
func (p *PlayerServiceImpl) AddPlayer(c *gin.Context, dto entity.AddPlayerDto) {
	_, err := dao.AddPlayer(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "新增失败")
	}
	result.Success(c, "新增成功")
}

// 获取详情
func (p *PlayerServiceImpl) GetPlayerById(c *gin.Context, id int) {
	player, err := dao.GetPlayerDetail(id)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "选手不存在")
	}
	result.Success(c, player)
}

var playerService = PlayerServiceImpl{}

func PlayerService() IPlayerService {
	return &playerService
}
