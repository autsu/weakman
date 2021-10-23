package dao

import (
	"github.com/sirupsen/logrus"
	"vote/errno"
	"vote/model"
	"vote/pkg/mysql"
)

type StuDao struct{}

func (d *StuDao) QueryByUsername(username string) (*model.Stu, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}

	var s model.Stu
	sql := `
select id, username, password, phone, name
from stu
where username = ?
`
	if err := db.Get(&s, sql, username); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectError.Error(), err)
		return nil, errno.MysqlSelectNoData
	}

	return &s, nil
}

func (d *StuDao) QueryByPhone(phone string) (*model.Stu, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}


	var s model.Stu
	sql := `
select id, username, password, phone, name
from stu
where phone = ?
`
	if err := db.Get(&s, sql, phone); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectError.Error(), err)
		return nil, errno.MysqlSelectNoData
	}

	return &s, nil
}

func (d *StuDao) QueryById(id string) (*model.Stu, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}

	sql := `
select id, username, password, phone, name
from stu
where id = ?
`
	var s model.Stu
	if err := db.Get(&s, sql, id); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectError.Error(), err)
		return nil, errno.MysqlSelectNoData
	}

	return &s, nil
}

func (d *StuDao) Insert(s *model.Stu) (n int64, err error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return 0, errno.MysqlConnectError
	}

	sql := `
insert into stu(username, password, phone, name)
values (?, ?, ?, ?)
`
	r, err := db.Exec(sql, &s.Username, &s.Password, &s.Phone, &s.Name)
	if err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlInsertError.Error(), err)
		return 0, errno.MysqlInsertError
	}
	n, _ = r.RowsAffected()
	return
}
