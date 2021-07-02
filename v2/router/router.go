package router

import (
	"github.com/gin-gonic/gin"
	"vote/v2/handler"
	"vote/v2/middleware"
)

func Router() *gin.Engine {
	var stuHandler handler.StuHandler
	var topicHandler handler.TopicHandler
	var topicOptionHandler handler.TopicOptionHandler
	var topicSetHandler handler.TopicSetHandler

	r := gin.Default()

	r.Use(middleware.Cors())

	r.POST("/topic", topicHandler.Insert)
	r.GET("/topic", topicHandler.QueryAllWithTopicSet)
	r.GET("/topic/:topicId", topicHandler.QueryById)
	r.GET("/topic/friendly", topicHandler.QueryAllFriendlyData)
	r.GET("/topic/friendly/:title", topicHandler.QueryByTitleFriendlyData)
	r.GET("/topic/result/:topicId", topicHandler.ShowResult)

	r.GET("/option/:topicId", topicOptionHandler.QueryByTopicId)

	r.GET("/topicset/:topicId", topicSetHandler.QueryByTopicId)
	r.POST("/topicset/vail/pwd", topicSetHandler.VailPassword)
	r.GET("/topic/participant/:optionId", topicOptionHandler.ShowParticipant)

	stu := r.Group("/stu")
	{
		stu.POST("/login", stuHandler.Login)
		stu.POST("/register", stuHandler.Register)

		r.Use(middleware.StuAuth)
		stu.GET("/info", stuHandler.GetInfo)
		stu.POST("/logout", stuHandler.Logout)
	}
	r.POST("/vote/single", topicOptionHandler.SingleVote)
	r.POST("/vote/multiple", topicOptionHandler.MultipleVote)
	return r
}
