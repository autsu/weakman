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
