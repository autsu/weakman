package service

import (
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"time"
	"vote/v2/dao"
	"vote/v2/enum"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
	"vote/v2/tool"
)

// TopicInsert
// 首先将数据插入到 vote_topic，返回插入的 id，再将其他数据插入到 topic_option 和 vote_set
func TopicInsert(t *model.Topic, s *model.TopicSet,
	o []*model.TopicOption, token string) error {

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

	friendlyData, err := toFriendlyDataWithSlice(topics)
	if err != nil {
		return nil, 0, err
	}

	count, err := dao.TopicCount()
	if err != nil {
		return nil, 0, err
	}

	logrus.Info("topic count: ", count)
	return friendlyData, count, nil
}

func TopicQueryByTitleFriendlyData(title string) (tf []*model.TopicFriendly, total int, err error) {
	topic, err := dao.TopicQueryByTitle(title)
	logrus.Infof("%+v\n", topic)
	if err != nil {
		return nil, 0, err
	}

	data, err := toFriendlyDataWithSlice(topic)
	if err != nil {
		return nil, 0, err
	}
	return data, len(data), nil
}

func toFriendlyDataWithSlice(topics []*model.Topic) ([]*model.TopicFriendly, error) {
	t := make([]*model.TopicFriendly, len(topics))
	// memset
	for i := 0; i < len(t); i++ {
		t[i] = &model.TopicFriendly{}
	}

	for i := 0; i < len(topics); i++ {
		t[i].Id = topics[i].Id
		logrus.Info("topics[i].StuId: ", topics[i].StuId)
		stu, err := dao.StuQueryById(topics[i].StuId)
		if err != nil {
			logrus.Error(errno.MysqlSelectError)
			return nil, err
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

	return t, nil
}

func toFriendlyData(topic *model.Topic) (*model.TopicFriendly, error) {
	var t model.TopicFriendly
	logrus.Infof("%+v\n", topic)
	t.Id = topic.Id
	t.Title = topic.Title
	logrus.Info("topics[i].StuId: ", topic.StuId)
	stu, err := dao.StuQueryById(topic.StuId)
	if err != nil {
		logrus.Error(errno.MysqlSelectError)
		return nil, err
	}
	t.StuName = stu.Name
	t.Title = topic.Title

	switch topic.TopicSets.SelectType {
	case enum.TOPIC_SELECT_TYPE_MULTIPLE_CHOICE:
		t.SelectType = "多选"
	case enum.TOPIC_SELECT_TYPE_SINGLE_CHOICE:
		t.SelectType = "单选"
	}

	switch topic.TopicSets.Anonymous {
	case enum.TOPIC_ANONYMOUS:
		t.Anonymous = "匿名"
	case enum.TOPIC_REAL_NAME:
		t.Anonymous = "实名"
	}

	if topic.TopicSets.Password == "" {
		t.NeedPassword = "公开"
	} else {
		t.NeedPassword = "私有"
	}

	switch topic.TopicSets.ShowResult {
	case enum.TOPIC_DONT_SHOW_RESULT:
		t.ShowResult = "不展示结果"
	case enum.TOPIC_SHOW_RESULT:
		t.ShowResult = "展示结果"
	}

	if time.Now().After(topic.Deadline) {
		t.Deadline = "进行中"
	} else {
		t.Deadline = "已结束"
	}
	return &t, nil
}

// IsVoted 用户是否已经投过票
func IsVoted(userId, topicId string) error {
	records, err := dao.RecordQueryByUserId(userId)
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, record := range records {
		option, _ := dao.TopicOptionQueryById(strconv.Itoa(record.OptionId))
		if option != nil && strconv.Itoa(option.TopicId) == topicId {
			return errno.TopicUserIsVoted
		}
	}
	return nil
}

func TopicQueryByIdWithFmtTime(topicId string) (*model.Topic, error) {
	topic, err := dao.TopicQueryById(topicId)
	if err != nil {
		return nil, err
	}
	topic.Deadline = tool.UtoB(topic.Deadline)
	topic.TopicSets.Password = "YOU DON'T NEED KNOWN"
	topic.CreateTime = tool.UtoB(topic.CreateTime)

	logrus.Info(topic)
	return topic, nil
}

func TopicShowResultById(id string) ([]*model.VoteResultVO, error) {
	rs, err := dao.TopicShowResultById(id)
	if err != nil {
		return nil, err
	}
	var total int
	for _, r := range rs {
		total += r.Votes
	}

	sort.Slice(rs, func(i, j int) bool {
		return rs[i].Votes > rs[j].Votes
	})

	for _, r := range rs {
		r.Percentage = float32(r.Votes) / float32(total)
	}

	return rs, nil
}