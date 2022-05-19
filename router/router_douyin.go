package router

import (
	"github.com/ACking-you/byte_douyin_project/handlers"
	"github.com/ACking-you/byte_douyin_project/models"
	"github.com/gin-gonic/gin"
)

func InitDouyinRouter() *gin.Engine {
	models.InitDB()
	r := gin.Default()

	r.Static("static", "./static")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", handlers.FeedVideoListHandler)

	apiRouter.GET("/user/", handlers.UserInfoHandler)

	apiRouter.POST("/user/login/", handlers.UserLoginHandler)
	apiRouter.POST("/user/register/", handlers.UserRegisterHandler)
	apiRouter.POST("/publish/action/", handlers.PublishVideoHandler)
	apiRouter.GET("/publish/list/", handlers.QueryVideoListHandler)
	return r
}
