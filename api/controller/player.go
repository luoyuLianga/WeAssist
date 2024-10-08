package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取选手列表
// @Tags      Player
// @Summary 获取选手列表接口
// @Produce json
// @Description 获取选手列表接口
// @Param data body entity.PlayerListDto true "data"
// @Success 200 {object} result.Result
// @router /api/player/list [post]
func GetPlayerList(c *gin.Context) {
	var dto entity.PlayerListDto
	_ = c.BindJSON(&dto)
	service.PlayerService().GetPlayerList(c, dto)
}

// 新增选手信息
// @Summary 更新选手信息接口
// @Tags      Player
// @Produce json
// @Description 更新选手信息接口
// @Param data body entity.AddPlayerDto true "data"
// @Success 200 {object} result.Result
// @router /api/player/add [post]
func AddPlayer(c *gin.Context) {
	var dto entity.AddPlayerDto
	_ = c.BindJSON(&dto)
	service.PlayerService().AddPlayer(c, dto)
}

// 获取选手详情
// @Summary 获取选手详情接口
// @Tags      Player
// @Produce json
// @Description 获取选手详情接口
// @Param id query int true "Id"
// @Success 200 {object} result.Result
// @router /api/player/:id [get]
func GetPlayerDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.PlayerService().GetPlayerById(c, id)
}

// 获取排行傍
// @Tags      Player
// @Summary 获取排行傍列表接口
// @Produce json
// @Description 获取排行傍列表接口
// @Param data body entity.PlayerListDto true "data"
// @Success 200 {object} result.Result
// @router /api/player/rank [post]
func GetRankList(c *gin.Context) {
	var dto entity.PlayerListDto
	_ = c.BindJSON(&dto)
	service.PlayerService().GetRankList(c, dto)
}
