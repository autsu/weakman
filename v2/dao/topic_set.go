package dao

import (
	"github.com/sirupsen/logrus"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
)

func TopicSetInsert(topicSet *model.TopicSet) (int64, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return -1, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
insert into vote_set(topic_id, select_type, anonymous, show_result, password) 
values (?, ?, ?, ?, ?)
`

	r, err := mysql.Exec(sql,
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

func TopicSetQueryByTopicId(topicId string) (*model.TopicSet, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select id, topic_id, select_type, anonymous, show_result, password
from vote_set
where topic_id = ?
`
	var ts model.TopicSet
	if err := mysql.Get(&ts, sql, topicId); err != nil {
		logrus.Warningf("%s: %s\n", errno.MysqlSelectNoData, err)
		return nil, errno.MysqlSelectNoData
	}
	return &ts, nil
}

func TopicSetQueryPasswordByTopicId(topicId string) (password string, err error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return "", errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select password
from vote_set
where topic_id = ?
`
	if err := mysql.Get(&password, sql, topicId); err != nil {
		logrus.Warningf("%s: %s\n", errno.MysqlSelectNoData, err)
		return "", errno.MysqlSelectNoData
	}
	return
}
