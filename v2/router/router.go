package router

import (
	"github.com/gin-gonic/gin"
	"vote/v2/handler"
	"vote/v2/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.POST("/topic", handler.TopicInsert)
	r.GET("/topic", handler.TopicQueryAllWithTopicSet)
	r.GET("/topic/:topicId", handler.TopicQueryById)
	r.GET("/topic/friendly", handler.TopicQueryAllFriendlyData)
	r.GET("/topic/friendly/:title", handler.TopicQueryByTitleFriendlyData)
	r.GET("/topic/result/:topicId", handler.TopicShowResult)


	r.GET("/option/:topicId", handler.TopicOptionQueryByTopicId)

	r.GET("/topicset/:topicId", handler.TopicSetQueryByTopicId)
	r.POST("/topicset/vail/pwd", handler.TopicSetVailPassword)


	stu := r.Group("/stu")
	{
		stu.POST("/login", handler.StuLogin)
		stu.POST("/register", handler.StuRegister)

		r.Use(middleware.StuAuth)
		stu.GET("/info", handler.StuGetInfo)
	}
	r.POST("/vote/single", handler.SingleVote)
	r.POST("/vote/multiple", handler.MultipleVote)
	return r
}
