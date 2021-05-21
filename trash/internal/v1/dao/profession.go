package dao

import (
	"log"
	model2 "vote/v1/model"
	pkg2 "vote/v1/pkg"
)

func ProfessionAdd(name, collegeId string) (int64, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sql := `
insert into stu_profession(college_id, name)
values (?, ?)
`

	r, err := mysql.Exec(sql, collegeId, name)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	n, _ := r.RowsAffected()
	return n, nil
}

func ProfessionQueryByName(name string) (*model2.Profession, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var p model2.Profession
	sql := `
select id,
       college_id,
       name
from stu_profession
where name = ?
  and delete_time IS NULL 
`
	if err := mysql.Get(&p, sql, name); err != nil {
		log.Println(err)
		return nil, err
	}
	return &p, nil
}

func ProfessionQueryById(id string) (*model2.Profession, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var p model2.Profession
	sql := `
select id,
       college_id,
       name
from stu_profession
where id = ?
  and delete_time IS NULL 
`
	if err := mysql.Get(&p, sql, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &p, nil
}

func ProfessionQueryByIdWithGrade(id string) (*model2.Profession, error) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var p model2.Profession
	sql := `
select sp.id, sp.college_id, sp.name, sg.id, sg.name
from stu_profession sp
         join stu_grade sg on sp.id = sg.profession_id
where sp.id = ?
`
	var gs []*model2.Grade
	rows, err := mysql.Query(sql, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var g model2.Grade
		if err := rows.Scan(
			&p.Id,
			&p.CollegeId,
			&p.Name,
			&g.Id,
			&g.Name); err != nil {
			log.Println(err)
			return nil, err
		} else {
			gs = append(gs, &g)
		}
	}

	p.Grades = gs
	return &p, nil
}
