package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func InitCrontab() {
	// 初始化 cron 实例
	c := cron.New()
	// 添加定时任务：每分钟执行一次  TODO 后面需要调整为每天0:05分开始执行
	_, err := c.AddFunc("* * * * *", func() {
		UseUserDayStats()
		ActiveUserDayStats()
	})
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}

	// 启动 cron 调度
	c.Start()
}
