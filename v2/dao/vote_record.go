package dao

import (
	"github.com/sirupsen/logrus"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
)

func RecordInsert(r *model.VoteRecord) (n int64, err error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return 0, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
insert into vote_record(uid, option_id, time) values (?, ?, ?)
`
	res, err := mysql.Exec(sql, r.Uid, r.OptionId, r.Time)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlInsertError, err)
		return 0, errno.MysqlInsertError
	}

	n, _ = res.RowsAffected()
	return n, nil
}

func RecordQueryByUserId(userId string) ([]*model.VoteRecord, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select id, uid, option_id, time  
from vote_record
where uid = ?
`
	var vs []*model.VoteRecord
	if err := mysql.Select(&vs, sql, userId); err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError, err)
		return nil, errno.MysqlSelectError
	}
	return vs, nil
}

func RecordQueryById(id string) ([]*model.VoteRecord, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select id, uid, option_id, time  
from vote_record
where id = ?
`
	var vs []*model.VoteRecord
	if err := mysql.Select(&vs, sql, id); err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError, err)
		return nil, errno.MysqlSelectError
	}
	return vs, nil
}
