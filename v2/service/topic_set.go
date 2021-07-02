package service

import (
	"vote/v2/dao"
	"vote/v2/errno"
	"vote/v2/model"
)

type TopicSetServer struct {
	topicSetDao dao.TopicSetDao
}

func (s *TopicSetServer) QueryByTopicId(topicId string) (*model.TopicSet, error) {
	return s.topicSetDao.QueryByTopicId(topicId)
}

func (s *TopicSetServer) VailPassword(inputPassword, topicId string) error {
	password, err := s.topicSetDao.QueryPasswordByTopicId(topicId)
	if err != nil {
		return err
	}
	if inputPassword != password {
		return errno.TopicPasswordIsWrong
	}
	return nil
}
