package service

import (
	"github.com/sirupsen/logrus"
	"time"
	"vote/v2/dao"
	"vote/v2/enum"
	"vote/v2/errno"
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
	// 转换为本地时间
	local := t.Deadline.Local()
	logrus.Info("local time: ", local)
	t.Deadline = local
	logrus.Info("t.Deadline: ", t.Deadline)
	if err := dao.TopicInsertWithSetAndOptions(t, s, o); err != nil {
		return err
	}

	return nil
}

func TopicQueryAllWithTopicSet(page, size int) ([]*model.Topic, error) {
	return dao.TopicQueryAllLimitWithTopicSet(page, size)
}

func TopicQueryAllFriendlyData(page, size int) (tf []*model.TopicFriendly, total int64, err error) {
	topics, err := dao.TopicQueryAllLimitWithTopicSet(page, size)
	if err != nil {
		return nil, 0, err
	}
	logrus.Info(topics)

	t := make([]*model.TopicFriendly, len(topics))
	// memset
	for i := 0; i < len(t); i++ {
		t[i] = &model.TopicFriendly{}
	}

	for i := 0; i < len(topics); i++ {
		t[i].Id = topics[i].Id
		logrus.Info("topics[i].StuId: "< topics[i].StuId)
		stu, err := dao.StuQueryById(topics[i].StuId)
		if err != nil {
			logrus.Error(errno.MysqlSelectError)
			break
		}
		t[i].StuName = stu.Name
		t[i].Title = topics[i].Title

		switch topics[i].TopicSets.SelectType {
		case enum.TOPIC_SELECT_TYPE_MULTIPLE_CHOICE:
			t[i].SelectType = "多选"
		case enum.TOPIC_SELECT_TYPE_SINGLE_CHOICE:
			t[i].SelectType = "单选"
		}

		switch topics[i].TopicSets.Anonymous {
		case enum.TOPIC_ANONYMOUS:
			t[i].Anonymous = "匿名"
		case enum.TOPIC_REAL_NAME:
			t[i].Anonymous = "实名"
		}

		if topics[i].TopicSets.Password == "" {
			t[i].NeedPassword = "公开"
		} else {
			t[i].NeedPassword = "私有"
		}

		switch topics[i].TopicSets.ShowResult {
		case enum.TOPIC_DONT_SHOW_RESULT:
			t[i].ShowResult = "不展示结果"
		case enum.TOPIC_SHOW_RESULT:
			t[i].ShowResult = "展示结果"
		}

		if time.Now().After(topics[i].Deadline) {
			t[i].Deadline = "进行中"
		} else {
			t[i].Deadline = "已结束"
		}
	}
	count, err := dao.TopicCount()
	if err != nil {
		return nil, 0, err
	}

	logrus.Info("topic count: ", count)
	return t, count, nil
}


