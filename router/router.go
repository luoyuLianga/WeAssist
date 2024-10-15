// 访问接口路由配置
package router

import (
	"WeAssist/api/controller"
	"WeAssist/common/config"
	"WeAssist/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 跌机是恢复
	router.Use(gin.Recovery())
	// 跨域中间件
	router.Use(middleware.Cors())
	// 图片访问路径静态文件夹可直接访问
	router.StaticFS(config.Config.ImageSettings.UploadDir,
		http.Dir(config.Config.ImageSettings.UploadDir))
	// 日志log中间件
	router.Use(middleware.Logger())
	// token验证中间件
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {
	// -----------------------------------swagger接口-----------------------------------
	// 放行swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
		user.GET("/get", controller.GetUser)
		user.PUT("/update", controller.UpdateUser)
	}
	userPlugin := router.Group("/api/userPlugin")
	{
		userPlugin.POST("/add", controller.AddUserPlugin)
	}
	qaRecord := router.Group("/api/qaRecord")
	{
		qaRecord.POST("/add", controller.AddQaRecord)
	}
	qaException := router.Group("/api/qaException")
	{
		qaException.POST("/add", controller.AddQaException)
	}
	operation := router.Group("/api/operation")
	{
		operation.POST("/add", controller.AddOperation)
		operation.GET("/get", controller.GetOperation)
		operation.PUT("/update", controller.UpdateOperation)
		operation.DELETE("/delete/:id", controller.DeleteOperation)
	}
	operationDayStats := router.Group("/api/operationDayStats")
	{
		operationDayStats.POST("/update", controller.UpdateOperationDayStats)
	}

	player := router.Group("/api/player")
	{
		player.POST("/list", controller.GetPlayerList)
		player.POST("/rank", controller.GetRankList)
		player.POST("/info", controller.GetPlayerDetail)
		player.POST("/add", controller.AddPlayer)
	}
	vote := router.Group("/api/vote")
	{
		vote.POST("/add", controller.AddVote)
	}
	activity := router.Group("/api/activity")
	{
		activity.POST("/add", controller.AddActivity)
	}
}
