package dao

import (
	"github.com/sirupsen/logrus"
	"strconv"
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

func TopicOptionAddNumber(i int32, optionId string) error {
	option, err := TopicOptionQueryById(optionId)
	if err != nil {
		return err
	}
	num32 := int32(option.Number)
	for atomic.CompareAndSwapInt32(&num32, num32, num32+i) {
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
	_, err = mysql.Exec(sql, num32, optionId)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlUpdateError, err)
		return errno.MysqlUpdateError
	}

	return nil
}

// TopicOptionSingleVote 封装了单选投票，事务操作
func TopicOptionSingleVote(r *model.VoteRecord, i int32, topicId string) error {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return errno.MysqlConnectError
	}
	defer mysql.Close()

	tx, err := mysql.Beginx()
	logrus.Info("事务开启")
	if err != nil {
		logrus.Errorf("%s: %v\n", errno.MysqlStartTransactionError, err)
		return errno.MysqlStartTransactionError
	}
	defer func() {
		if p := recover(); p != nil {
			logrus.Errorf("事务回滚")
			tx.Rollback()
			// panic(p) // re-throw panic after Rollback
		} else if err != nil {
			logrus.Errorf("事务回滚")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			logrus.Info("事务提交")
		}
	}()

	// 1.插入数据到投票记录表
	sql := `
insert into vote_record(uid, option_id, time) values (?, ?, ?)
`
	_, err = tx.Exec(sql, r.Uid, r.OptionId, r.Time)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlInsertError, err)
		return errno.MysqlInsertError
	}

	// 2. 根据 option_id 查询出该选项的票数
	//该选项的票数 + i
	sql1 := `
select id,
       topic_id,
       option_content,
       number
from topic_option
where id = ?
`
	var o model.TopicOption
	if err := tx.Get(&o, sql1, strconv.Itoa(r.OptionId)); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectNoData, err)
		return errno.MysqlSelectNoData
	}
	//option, err := TopicOptionQueryById(strconv.Itoa(r.OptionId))
	//if err != nil {
	//	return err
	//}

	// 3. 票数原子+i，并将新数据写入
	num32 := int32(o.Number)
	for atomic.CompareAndSwapInt32(&num32, num32, num32+i) {
		break
	}

	sql2 := `
update topic_option set number = ? 
where id = ?
`
	_, err = tx.Exec(sql2, num32, r.OptionId)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlUpdateError, err)
		return errno.MysqlUpdateError
	}

	return nil
}

// TopicOptionMultipleVote 封装了多选投票，事务操作
func TopicOptionMultipleVote(rs []*model.VoteRecord, i int32, topicId string) error {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return errno.MysqlConnectError
	}
	defer mysql.Close()

	tx, err := mysql.Beginx()
	logrus.Info("事务开启")
	if err != nil {
		logrus.Errorf("%s: %v\n", errno.MysqlStartTransactionError, err)
		return errno.MysqlStartTransactionError
	}
	defer func() {
		if p := recover(); p != nil {
			logrus.Errorf("事务回滚")
			tx.Rollback()
			// panic(p) // re-throw panic after Rollback
		} else if err != nil {
			logrus.Errorf("事务回滚")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			logrus.Info("事务提交")
		}
	}()

	// FIXME for 循环里抛出 error 不会执行到上面的 defer 里
	for _, r := range rs {
		// 1.插入数据到投票记录表
		sql := `
insert into vote_record(uid, option_id) values (?, ?)
`
		_, err = tx.Exec(sql, r.Uid, r.OptionId)
		if err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlInsertError, err)
			logrus.Errorf("事务回滚")
			tx.Rollback()
			return errno.MysqlInsertError
		}

		// 2. 根据 option_id 查询出该选项的票数
		sql1 := `
select id,
       topic_id,
       option_content,
       number
from topic_option
where id = ?
`
		var o model.TopicOption
		if err := tx.Get(&o, sql1, strconv.Itoa(r.OptionId)); err != nil {
			logrus.Warnf("%s: %s\n", errno.MysqlSelectNoData, err)
			logrus.Errorf("事务回滚")
			tx.Rollback()
			return errno.MysqlSelectNoData
		}

		num32 := int32(o.Number)
		for atomic.CompareAndSwapInt32(&num32, num32, num32+i) {
			break
		}

		// 3. 票数原子+i，并将新数据写入
		sql2 := `
update topic_option set number = ? 
where id = ?
`
		_, err = tx.Exec(sql2, num32, r.OptionId)
		if err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlUpdateError, err)
			logrus.Errorf("事务回滚")
			tx.Rollback()
			return errno.MysqlUpdateError
		}
	}

	return nil
}

