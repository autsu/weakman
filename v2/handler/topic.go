package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"vote/v2/enum/result"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/service"
)

func TopicInsert(c *gin.Context) {
	var topicVo = struct {
		// topic
		Title       string
		Description string
		Deadline    time.Time

		// topic_set
		SelectType int
		Anonymous  int
		ShowResult int
		Password   string

		// topic_option
		Option []struct {
			OptionContent string
		}
	}{}

	token := c.GetHeader("Authorization")
	if err := c.ShouldBindJSON(&topicVo); err != nil {
		logrus.Error("bind json to model.topic error: ", err)
		c.JSON(http.StatusBadRequest, result.NewWithCode(result.BAD_REQUEST))
		return
	}
	logrus.Infof("topicVo: %+v\n", topicVo)

	t := &model.Topic{
		Title:       topicVo.Title,
		Description: topicVo.Description,
		Deadline:    topicVo.Deadline,
	}

	s := &model.TopicSet{
		SelectType: topicVo.SelectType,
		Anonymous:  topicVo.Anonymous,
		ShowResult: topicVo.ShowResult,
		Password:   topicVo.Password,
	}

	var o []*model.TopicOption
	for _, v := range topicVo.Option {
		op := &model.TopicOption{
			OptionContent: v.OptionContent,
		}
		o = append(o, op)
	}

	if err := service.TopicInsert(t, s, o, token); err != nil {
		if errors.Is(err, errno.TokenInvalid) {
			logrus.Error(err)
			c.JSON(http.StatusOK, result.NewWithCode(result.TOKEN_INVALID))
			return
		}
		c.JSON(http.StatusOK, result.NewWithCode(result.SERVER_ERROR))
		return
	}

	c.JSON(http.StatusOK,
		result.NewWithCodeAndData(result.SUCCESS, nil))
}
