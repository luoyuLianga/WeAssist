package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// 投票
func AddVote(dto entity.VoteDto) error {
	user := entity.Vote{
		UserId:     dto.UserId,
		PlayerId:   dto.PlayerId,
		CreateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&user).Error
	return err
}

// 获取投票信息
func GetVoteInfo(dto entity.VoteDto) (vote entity.Vote, err error) {
	err = db.Db.Where("user_id = ? AND player_id = ?", dto.UserId, dto.PlayerId).First(&vote).Error
	return vote, err
}
