package model

type Stu struct {
	Id       int
	Username string
	Password string
	Phone    string
	ClassId  int `db:"class_id"`
	Name     string
	Job      int
}

type College struct {
	Id          int
	Name        string
	Professions []*Profession
}

type Profession struct {
	Id        int
	CollegeId int `db:"college_id"`
	Name      string
	Grades    []*Grade
}

type Grade struct {
	Id           int
	ProfessionId int `db:"profession_id"`
	Name         string
}

type Class struct {
	Id      int
	GradeId int `db:"grade_id"`
	Name    string
}

type Admin struct {
	Id       string
	Username string
	Password string
	Name     string
}
