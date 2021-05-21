package service

import (
	"vote/v2/dao"
	"vote/v2/model"
	"vote/v2/pkg"
)

// TopicInsert
// 首先将数据插入到 vote_topic，返回插入的 id，再将其他数据插入到 topic_option 和 vote_set
func TopicInsert(
	t *model.Topic,
	s *model.TopicSet,
	o []*model.TopicOption,
	token string) error {

	tok, err := pkg.ParseTokenWithBearer(token)
	if err != nil {
		return err
	}

	t.StuId = tok.Id
	if err := dao.TopicInsertWithSetAndOptions(t, s, o); err != nil {
		return err
	}

	//lastId, err := dao.TopicInsert(t)
	//if err != nil {
	//	return err
	//}
	//
	//
	//_, err = dao.TopicSetInsert(s)
	//if err != nil {
	//	return err
	//}
	//
	//for _, option := range o {
	//	option.TopicId = int(lastId)
	//	_, err = dao.TopicOptionInsert(o)
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}


