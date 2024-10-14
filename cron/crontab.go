package cron

import (
	"WeAssist/common/config"
	"WeAssist/pkg/log"
	"fmt"
	"github.com/robfig/cron/v3"
)

func InitCrontab() {
	// 初始化 cron 实例
	c := cron.New()

	log.Log().Infof("cron run time:%s", config.Config.Cron.Run)
	// 添加定时任务：每分钟执行一次  TODO 后面需要调整为每天0:05分开始执行
	_, err := c.AddFunc(config.Config.Cron.Run, func() {

		// 用户数据统计
		{
			//UseUserDayStats()
			//ActiveUserDayStats()
		}

		// QA数据统计
		{
			QaDayStats()
		}

	})
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}

	// 启动 cron 调度
	c.Start()
}
