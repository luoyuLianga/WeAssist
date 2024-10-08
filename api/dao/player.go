package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"

	"gorm.io/gorm"
)

// 获取选手列表
func GetPlayerList(aid uint, sort string) ([]entity.Player, error) {
	var players []entity.Player
	err := db.Db.Where("aid = ?", aid).Order(sort).Find(&players).Error
	return players, err
}

// 更新选手
func UpdatePlayer(id int) (player entity.Player, err error) {
	err = db.Db.Model(&player).Where("id = ?", id).UpdateColumn("update_time", util.HTime{Time: time.Now()}).UpdateColumn("score", gorm.Expr("score +?", 1)).Error
	return player, err
}

// 创建选手
func AddPlayer(dto entity.AddPlayerDto) (id uint, err error) {
	player := entity.Player{
		Aid:         dto.Aid,
		Ref:         dto.Ref,
		Nickname:    dto.Nickname,
		Declaration: dto.Declaration,
		Avatar:      dto.Avatar,
		Score:       0,
		Phone:       dto.Phone,
		Note:        dto.Note,
		CreateTime:  util.HTime{Time: time.Now()},
		UpdateTime:  util.HTime{Time: time.Now()},
	}
	err = db.Db.Create(&player).Error
	return player.ID, err
}

// 获取选手详情
func GetPlayerDetail(aid int) (entity.Player, error) {
	var player entity.Player
	err := db.Db.Where("id = ?", aid).First(&player).Error
	return player, err
}

// 更具ids获取选手列表
func GetPlayerByIds(aids []int) (players []entity.Player, err error) {
	err = db.Db.Where("id in (?)", aids).Order("score desc").Find(&players).Error
	return players, err
}
