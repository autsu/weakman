package dao

import (
	"log"
	"vote/v2/model"
	"vote/v2/pkg"
)

func AdminQueryByUname(username string) (*model.Admin, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
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
	if err := mysql.Get(&a, sql, username); err != nil {
		log.Println(err)
		return nil, err
	}
	return &a, nil
}
