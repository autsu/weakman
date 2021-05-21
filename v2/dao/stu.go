package dao

import (
	"github.com/sirupsen/logrus"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
)

func StuQueryByUsername(username string) (*model.Stu, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	var s model.Stu
	sql := `
select id, username, password, phone, name
from stu
where username = ?
`
	if err := mysql.Get(&s, sql, username); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectError.Error(), err)
		return nil, errno.MysqlSelectNoData
	}

	return &s, nil
}

func StuQueryByPhone(phone string) (*model.Stu, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	var s model.Stu
	sql := `
select id, username, password, phone, name
from stu
where phone = ?
`
	if err := mysql.Get(&s, sql, phone); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectError.Error(), err)
		return nil, errno.MysqlSelectNoData
	}

	return &s, nil
}

func StuQueryById(id string) (*model.Stu, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlSelectError
	}
	defer mysql.Close()

	sql := `
select id, username, password, phone, name
from stu
where id = ?
`
	var s model.Stu
	if err := mysql.Get(&s, sql, id); err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlSelectError.Error(), err)
		return nil, errno.MysqlSelectNoData
	}

	return &s, nil
}

func StuInsert(s *model.Stu) (n int64, err error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return 0, errno.MysqlSelectError
	}
	defer mysql.Close()

	sql := `
insert into stu(username, password, phone, name)
values (?, ?, ?, ?)
`
	r, err := mysql.Exec(sql, &s.Username, &s.Password, &s.Phone, &s.Name)
	if err != nil {
		logrus.Warnf("%s: %s\n", errno.MysqlInsertError.Error(), err)
		return 0, errno.MysqlInsertError
	}
	n, _ = r.RowsAffected()
	return
}


