package dao

import (
	"vote/errno"
	"vote/model"
	"vote/pkg/mysql"

	"github.com/sirupsen/logrus"
)

type TopicSetDao struct{}

func (d *TopicSetDao) Insert(topicSet *model.TopicSet) (int64, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return 0, errno.MysqlConnectError
	}

	sql := `
insert into vote_set(topic_id, select_type, anonymous, show_result, password) 
values (?, ?, ?, ?, ?)
`

	r, err := db.Exec(sql,
		topicSet.TopicId,
		topicSet.SelectType,
		topicSet.Anonymous,
		topicSet.ShowResult,
		topicSet.Password)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlInsertError, err)
		return -1, errno.MysqlInsertError
	}

	n, _ := r.RowsAffected()

	return n, nil
}

func (d *TopicSetDao) QueryByTopicId(topicId string) (*model.TopicSet, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}

	sql := `
select id, topic_id, select_type, anonymous, show_result, password
from vote_set
where topic_id = ?
`
	var ts model.TopicSet
	if err := db.Get(&ts, sql, topicId); err != nil {
		logrus.Warningf("%s: %s\n", errno.MysqlSelectNoData, err)
		return nil, errno.MysqlSelectNoData
	}
	return &ts, nil
}

func (d *TopicSetDao) QueryPasswordByTopicId(topicId string) (password string, err error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return "", errno.MysqlConnectError
	}

	sql := `
select password
from vote_set
where topic_id = ?
`
	if err := db.Get(&password, sql, topicId); err != nil {
		logrus.Warningf("%s: %s\n", errno.MysqlSelectNoData, err)
		return "", errno.MysqlSelectNoData
	}
	return
}
