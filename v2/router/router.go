package router

import (
	"github.com/gin-gonic/gin"
	"vote/v2/handler"
	"vote/v2/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	stu := r.Group("/stu")
	{
		stu.POST("/login", handler.StuLogin)
		stu.POST("/register", handler.StuRegister)

		r.Use(middleware.StuAuth)
		stu.GET("/info", handler.StuGetInfo)
	}

	r.POST("/topic", handler.TopicInsert)
	r.GET("/topic", handler.TopicQueryAllWithTopicSet)
	r.GET("/topic/list", handler.TopicQueryAllFriendlyData)

	return r
}
