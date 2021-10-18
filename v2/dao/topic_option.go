package dao

import (
	"strconv"
	"sync/atomic"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
	"github.com/sirupsen/logrus"
)

type TopicOptionDao struct {}

func (d *TopicOptionDao) Insert(o []*model.TopicOption) (int64, error) {
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

func (d *TopicOptionDao) QueryByTopicId(topicId string) (ops []*model.TopicOption, total int, err error) {
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

func (d *TopicOptionDao) QueryById(id string) (*model.TopicOption, error) {
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

func (d *TopicOptionDao) AddNumber(i int32, optionId string) error {
	option, err := d.QueryById(optionId)
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

// SingleVote 封装了单选投票，事务操作
func (d *TopicOptionDao) SingleVote(r *model.VoteRecord, i int32, topicId string) error {
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

	// 2. 根据 option_id 查询出该选项的票数，然后将该选项的票数 + i，并写入数据库
	//
	// 这里的 sql 语句需要添加 for update，为 select 添加排他锁，在当前事务
	// 未提交之前，其他所有事务都会在此被阻塞
	// 如果这里不加锁，那么可能会出现更新覆盖的问题，比如 A、B 2 条线程同时进来，
	// A 得到选项1 的值为 100，但在 A 将值更新为 100 + 1 之前，
	// B 也得到了选项1 的值 100，之后 A 更新为 101，B 也更新为 100 + 1，
	// 导致最终结果为 101 而不是正确的 102，因为 B 把 A 的更新给覆盖了
	//
	// 在数据库层面加了排它锁以后，就不需要在应用层面，对票数 + 1 这一步进行原子
	// 操作了 
	sql1 := `
select id,
       topic_id,
       option_content,
       number
from topic_option
where id = ? for update
`
	var o model.TopicOption
	if err := tx.Get(&o, sql1, strconv.Itoa(r.OptionId)); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectNoData, err)
		return errno.MysqlSelectNoData
	}
	logrus.Println("old: ", o.Number)

	// 3. 票数原子+i，并将新数据写入
	// num32 := int32(o.Number)
	// for atomic.CompareAndSwapInt32(&num32, num32, num32+i) {
	// 	break
	// }
	// logrus.Println("old+1: ", num32)

	sql2 := `
update topic_option set number = ? 
where id = ?;
`
	_, err = tx.Exec(sql2, o.Number+1, r.OptionId)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlUpdateError, err)
		return errno.MysqlUpdateError
	}

	return nil
}

// MultipleVote 封装了多选投票，事务操作
func (d *TopicOptionDao) MultipleVote(rs []*model.VoteRecord, i int32, topicId string) error {
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
where id = ? for update
`
		var o model.TopicOption
		if err := tx.Get(&o, sql1, strconv.Itoa(r.OptionId)); err != nil {
			logrus.Warnf("%s: %s\n", errno.MysqlSelectNoData, err)
			logrus.Errorf("事务回滚")
			tx.Rollback()
			return errno.MysqlSelectNoData
		}

		// 加了 for update 不需要原子操作了
		// num32 := int32(o.Number)
		// for atomic.CompareAndSwapInt32(&num32, num32, num32+i) {
		// 	break
		// }

		// 3. 票数原子+i，并将新数据写入
		sql2 := `
update topic_option set number = ? 
where id = ?
`
		_, err = tx.Exec(sql2, o.Number+1, r.OptionId)
		if err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlUpdateError, err)
			logrus.Errorf("事务回滚")
			tx.Rollback()
			return errno.MysqlUpdateError
		}
	}

	return nil
}

// ShowParticipantById 查看该选项的参与者
func (d *TopicOptionDao) ShowParticipantById(id string) (participantName []string, err error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select s.name from vote_record vr 
join stu s on s.id = vr.uid 
where option_id = ?
`

	if err := mysql.Select(&participantName, sql, id); err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError)
		return nil, errno.MysqlSelectError
	}

	return participantName, nil
}
