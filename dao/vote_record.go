package dao

import (
	"vote/errno"
	"vote/model"
	"vote/pkg/mysql"

	"github.com/sirupsen/logrus"
)

type VoteRecordDao struct{}

func (d *VoteRecordDao) Insert(r *model.VoteRecord) (n int64, err error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return 0, errno.MysqlConnectError
	}

	sql := `
insert into vote_record(uid, option_id, time) values (?, ?, ?)
`
	res, err := db.Exec(sql, r.Uid, r.OptionId, r.Time)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlInsertError, err)
		return 0, errno.MysqlInsertError
	}

	n, _ = res.RowsAffected()
	return n, nil
}

func (d *VoteRecordDao) QueryByUserId(userId string) ([]*model.VoteRecord, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}

	sql := `
select id, uid, option_id, time  
from vote_record
where uid = ?
`
	var vs []*model.VoteRecord
	if err := db.Select(&vs, sql, userId); err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError, err)
		return nil, errno.MysqlSelectError
	}
	return vs, nil
}

func (d *VoteRecordDao) QueryById(id string) ([]*model.VoteRecord, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}

	sql := `
select id, uid, option_id, time  
from vote_record
where id = ?
`
	var vs []*model.VoteRecord
	if err := db.Select(&vs, sql, id); err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError, err)
		return nil, errno.MysqlSelectError
	}
	return vs, nil
}
