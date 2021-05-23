package service

import (
	"vote/v2/dao"
	"vote/v2/model"
)

func TopicSetQueryByTopicId(topicId string) (*model.TopicSet, error) {
	return dao.TopicSetQueryByTopicId(topicId)
}
