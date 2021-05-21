package dao

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
	model2 "vote/v1/model"
	pkg2 "vote/v1/pkg"
)

func CollegeAdd(name string) (int64, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sql := `
insert into stu_college(name)
values (?)
`

	r, err := mysql.Exec(sql, name)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	n, _ := r.RowsAffected()
	return n, nil
}

func CollegeUpdateName(oldName, newName string) (int64, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sql := `
update stu_college
set name = 'test1'
where name = 'test'
`

	r, err := mysql.Exec(sql, newName, oldName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	n, _ := r.RowsAffected()
	return n, nil
}

// CollegeQueryByName An error is returned if the result set is empty.
func CollegeQueryByName(name string) (*model2.College, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var c model2.College
	sql := `
select id,
       name
from stu_college
where name = ?
  and delete_time IS NULL 
`
	if err := mysql.Get(&c, sql, name); err != nil {
		log.Println(err)
		return nil, err
	}
	return &c, nil
}

func CollegeQueryById(id string) (*model2.College, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var c model2.College
	sql := `
select id,
       name
from stu_college
where id = ?
	and delete_time IS NULL 
`
	if err := mysql.Get(&c, sql, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &c, nil
}

func CollegeQueryByIdWithProfession(id string) (*model2.College, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sql := `
select sc.id as id, sc.name as name, sp.id, sp.name
from stu_college sc
         join stu_profession sp on sc.id = sp.college_id
where sc.id = ?
`
	var c model2.College

	var ps []*model2.Profession
	rows, err := mysql.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p model2.Profession
		if err := rows.Scan(&c.Id, &c.Name, &p.Id, &p.Name); err != nil {
			log.Println(err)
			return nil, err
		} else {
			ps = append(ps, &p)
		}
	}

	c.Professions = ps
	return &c, nil
}

func CollegeSoftDelete(id string) (int64, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sql := `
update stu_college
set delete_time = ?
where id = ?
`
	r, err := mysql.Exec(sql, time.Now(), id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	n, _ := r.RowsAffected()
	return n, nil
}

func CollegeSoftDeleteMore(ids []string) (n int64, err error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sql := `
update stu_college
set delete_time = ?
where id in (?)
`
	query, args, err := sqlx.In(sql, time.Now(), interface{}(ids))
	if err != nil {
		log.Println(err)
		return 0, err
	}
	//log.Println(query)

	r, err := mysql.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	n, _ = r.RowsAffected()
	return
}

func CollegeQueryAll() ([]*model2.College, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	sql := `
select id, name
from stu_college
where delete_time is NULL
`
	var cs []*model2.College
	if err := mysql.Select(&cs, sql); err != nil {
		log.Println(err)
		return nil, err
	}
	return cs, nil
}

func CollegeUpdate(coll *model2.College) (n int64, err error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sql := `
update stu_college
set name = ?
where id = ?
`

	r, err := mysql.Exec(sql, coll.Name, coll.Id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	n, _ = r.RowsAffected()
	return
}
