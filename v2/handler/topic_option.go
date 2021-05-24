package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"vote/v2/enum/result"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/service"
)

func TopicOptionQueryByTopicId(c *gin.Context) {
	topicId := c.Param("topicId")
	options, total, err := service.TopicOptionQueryByTopicId(topicId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCodeAndData(result.SUCCESS, gin.H{
		"option": options,
		"total": total,
	}))
}

func SingleVote(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var v model.VoteSingleVO
	if err := c.ShouldBindJSON(&v); err != nil {
		logrus.Error("bind json to model.topic error: ", err)
		c.JSON(http.StatusBadRequest, result.NewWithCode(result.BAD_REQUEST))
		return
	}

	if err := service.SingleVote(v.Record, v.Votes, strconv.Itoa(v.TopicId), token); err != nil {
		if errors.Is(err, errno.TopicUserIsVoted) {
			c.JSON(http.StatusOK, result.NewWithCode(result.USER_IS_VOTED))
			return
		}
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCode(result.SUCCESS))
}

func MultipleVote(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var v model.VoteMultipleVO
	if err := c.ShouldBindJSON(&v); err != nil {
		logrus.Error("bind json to model.topic error: ", err)
		c.JSON(http.StatusBadRequest, result.NewWithCode(result.BAD_REQUEST))
		return
	}

	if err := service.MultipleVote(v.Record, v.Votes, strconv.Itoa(v.TopicId), token); err != nil {
		if errors.Is(err, errno.TopicUserIsVoted) {
			c.JSON(http.StatusOK, result.NewWithCode(result.USER_IS_VOTED))
			return
		}
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCode(result.SUCCESS))
}