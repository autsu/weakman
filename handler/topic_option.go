package handler

import (
	"errors"
	"net/http"
	"strconv"
	"vote/enum/result"
	"vote/errno"
	"vote/model"
	"vote/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TopicOptionHandler struct {
	topicOptionService service.TopicOptionService
}

func (h *TopicOptionHandler) QueryByTopicId(c *gin.Context) {
	topicId := c.Param("topicId")
	options, total, err := h.topicOptionService.QueryByTopicId(topicId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCodeAndData(result.SUCCESS, gin.H{
		"option": options,
		"total":  total,
	}))
}

func (h *TopicOptionHandler) SingleVote(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var v model.VoteSingleVO
	if err := c.ShouldBindJSON(&v); err != nil {
		logrus.Error("bind json to model.topic error: ", err)
		c.JSON(http.StatusBadRequest, result.NewWithCode(result.BAD_REQUEST))
		return
	}

	if err := h.topicOptionService.SingleVote(v.Record, v.Votes, strconv.Itoa(v.TopicId), token); err != nil {
		if errors.Is(err, errno.TopicUserIsVoted) {
			c.JSON(http.StatusOK, result.NewWithCode(result.USER_IS_VOTED))
			return
		}
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCode(result.SUCCESS))
}

func (h *TopicOptionHandler) MultipleVote(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var v model.VoteMultipleVO
	if err := c.ShouldBindJSON(&v); err != nil {
		logrus.Error("bind json to model.topic error: ", err)
		c.JSON(http.StatusBadRequest, result.NewWithCode(result.BAD_REQUEST))
		return
	}

	if err := h.topicOptionService.MultipleVote(v.Record, v.Votes, strconv.Itoa(v.TopicId), token); err != nil {
		if errors.Is(err, errno.TopicUserIsVoted) {
			c.JSON(http.StatusOK, result.NewWithCode(result.USER_IS_VOTED))
			return
		}
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCode(result.SUCCESS))
}

func (h *TopicOptionHandler) ShowParticipant(c *gin.Context) {
	optionId := c.Param("optionId")
	participant, err := h.topicOptionService.ShowParticipant(optionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.NewWithCode(result.SERVER_ERROR))
		return
	}
	c.JSON(http.StatusOK, result.NewWithCodeAndData(result.SUCCESS, participant))
}
