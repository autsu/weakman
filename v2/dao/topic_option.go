package dao

import (
	"github.com/sirupsen/logrus"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
)

func TopicOptionInsert(o []*model.TopicOption) (int64, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return -1, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
insert into topic_option(topic_id, option_content) 
values (?, ?)
`

	var n int64
	for _, option := range o {
		r, err := mysql.Exec(sql, option.TopicId, option.OptionContent)
		if err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlInsertError, err)
			return -1, errno.MysqlInsertError
		}
		nn, _ := r.RowsAffected()
		n += nn
	}

	return n, nil
}