package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"vote/v2/enum/result"
	"vote/v2/errno"
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

func TopicSetVailPassword(c *gin.Context) {
	var s = struct {
		TopicId  int	`json:"topic_id"`
		Password string	`json:"password"`
	}{}
	if err := c.ShouldBindJSON(&s); err != nil {
		logrus.Error("bind json to model.topic error: ", err)
		c.JSON(http.StatusBadRequest, result.NewWithCode(result.BAD_REQUEST))
		return
	}

	if err := service.TopicSetVailPassword(s.Password, strconv.Itoa(s.TopicId)); err != nil {
		if errors.Is(err, errno.TopicPasswordIsWrong) {
			c.JSON(http.StatusOK, result.NewWithCode(result.TOPIC_PASSWORD_ERROR))
			return
		}
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}

	c.JSON(http.StatusOK, result.NewWithCodeAndData(result.SUCCESS, "ok"))
}
