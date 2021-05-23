package dao

import (
	"github.com/sirupsen/logrus"
	"sync/atomic"
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

func TopicOptionQueryByTopicId(topicId string) (ops []*model.TopicOption, total int, err error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, 0, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select id,
       topic_id,
       option_content,
       number
from topic_option
where topic_id = ?
`
	var os []*model.TopicOption
	if err := mysql.Select(&os, sql, topicId); err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError, err)
		return nil, 0, errno.MysqlSelectError
	}

	return os, len(os), nil
}

func TopicOptionQueryById(id string) (*model.TopicOption, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select id,
       topic_id,
       option_content,
       number
from topic_option
where id = ?
`
	var o model.TopicOption
	if err := mysql.Get(&o, sql, id); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectNoData, err)
		return nil, errno.MysqlSelectNoData
	}

	return &o, err
}

// TODO 添加事务
func TopicOptionAddNumber(i int32, optionId string) error {
	option, err := TopicOptionQueryById(optionId)
	if err != nil {
		return err
	}
	id32 := int32(option.Id)
	for atomic.CompareAndSwapInt32(&id32, id32, id32+i) {
		break
	}

	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
update topic_option set number = ? 
where id = ?
`
	_, err = mysql.Exec(sql, id32, optionId)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlUpdateError, err)
		return errno.MysqlUpdateError
	}

	return nil
}
