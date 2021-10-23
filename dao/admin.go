package dao

import (
	"log"
	"vote/errno"
	"vote/model"
	"vote/pkg/mysql"
	"github.com/sirupsen/logrus"
)

func AdminQueryByUname(username string) (*model.Admin, error) {
	db, err := mysql.GetConn()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError.Error(), err)
		return nil, errno.MysqlConnectError
	}

	sql := `
select id,
       username,
       password
from admin
where username = ?
  and delete_time is null
`
	var a model.Admin
	if err := db.Get(&a, sql, username); err != nil {
		log.Println(err)
		return nil, err
	}
	return &a, nil
}
