package model

import (
	"encoding/json"
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
	StuId              string `db:"stu_id"`
	Title, Description string
	Deadline           time.Time // json 需要绑定为 UTC 时间格式，例如：020-07-31T14:27:10.035542+08:00
	ReviewStatus       int       `db:"review_status"`
	CreateTime         time.Time `db:"create_time"`
	//DeleteTime         time.Time `db:"delete_time"`
	TopicSets          *TopicSet
}

type TopicSet struct {
	Id         int
	TopicId    int `db:"topic_id"`
	SelectType int `db:"select_type"`
	Anonymous  int
	ShowResult int `db:"show_result"`
	Password   string
}

type TopicOption struct {
	Id            int
	TopicId       int
	OptionContent string
	Number        int
}

// BJT 北京时间 unused
type BJT struct {
	Deadline time.Time
}

func (b *BJT) MarshalJSON() ([]byte, error) {
	type alias BJT

	return json.Marshal(struct {
		*alias
		Deadline string
	}{
		alias:    (*alias)(b),
		Deadline: b.Deadline.Format("2006-01-02 15:04:05"),
	})
}

func (b *BJT) UnmarshalJSON(data []byte) error {
	type alias BJT

	tmp := &struct {
		*alias
		Deadline string
	}{
		alias: (*alias)(b),
	}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	b.Deadline, err = time.Parse(`2006-01-02 15:04:05`, tmp.Deadline)
	if err != nil {
		return err
	}

	return nil
}
