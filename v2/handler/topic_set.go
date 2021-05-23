package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vote/v2/enum/result"
	"vote/v2/service"
)

func TopicSetQueryByTopicId(c *gin.Context) {
	topicId := c.Param("topicId")
	topicSet, err := service.TopicSetQueryByTopicId(topicId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCodeAndData(result.SUCCESS, topicSet))
}
