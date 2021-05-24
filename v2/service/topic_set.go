package service

import (
	"vote/v2/dao"
	"vote/v2/errno"
	"vote/v2/model"
)

func TopicSetQueryByTopicId(topicId string) (*model.TopicSet, error) {
	return dao.TopicSetQueryByTopicId(topicId)
}

func TopicSetVailPassword(inputPassword, topicId string) error {
	password, err := dao.TopicSetQueryPasswordByTopicId(topicId)
	if err != nil {
		return err
	}
	if inputPassword != password {
		return errno.TopicPasswordIsWrong
	}
	return nil
}
