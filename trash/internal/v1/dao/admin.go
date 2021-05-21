package dao

import (
	"log"
	model2 "vote/v1/model"
	pkg2 "vote/v1/pkg"
)

func AdminQueryByUname(username string) (*model2.Admin, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sql := `
select id,
       username,
       password
from vote_admin
where username = ?
  and delete_time is null
`
	var a model2.Admin
	if err := mysql.Get(&a, sql, username); err != nil {
		log.Println(err)
		return nil, err
	}
	return &a, nil
}
