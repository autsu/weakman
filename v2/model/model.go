package model

import (
	"time"
)

type Stu struct {
	Id       int
	Username string
	Password string
	Phone    string
	Name     string
}

type Topic struct {
	Id                 int
	StuId              string
	Title, Description string
	// json 绑定格式：2020-07-31T14:27:10.035542+08:00
	Deadline time.Time
	Status   int
}

type TopicSet struct {
	Id         int
	TopicId    int
	SelectType int
	Anonymous  int
	ShowResult int
	Password   string
}

type TopicOption struct {
	Id            int
	TopicId       int
	OptionContent string
	Number        int
}
