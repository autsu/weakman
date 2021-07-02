package dao

import (
	"github.com/sirupsen/logrus"
	"vote/v2/enum"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
)

type TopicDao struct {
}

func (d *TopicDao) Insert(t *model.Topic) (id int64, err error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return -1, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
insert into vote_topic(stu_id, title, description, deadline)  
values (?, ?, ?, ?)
`

	r, err := mysql.Exec(sql, t.StuId, t.Title, t.Description, t.Deadline)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlInsertError, err)
		return -1, errno.MysqlInsertError
	}

	lastId, err := r.LastInsertId()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlGetLastInsertIdError, err)
		return -1, errno.MysqlGetLastInsertIdError
	}
	return lastId, err
}

// InsertWithSetAndOptions 因为要连续插入到三张表，所以需要添加事务
func (d *TopicDao) InsertWithSetAndOptions(
	t *model.Topic,
	s *model.TopicSet,
	o []*model.TopicOption) error {

	logrus.Info(t)
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return errno.MysqlConnectError
	}
	defer mysql.Close()

	tx, err := mysql.Beginx()
	logrus.Info("事务开启")
	if err != nil {
		logrus.Errorf("%s: %v\n", errno.MysqlStartTransactionError, err)
		return errno.MysqlStartTransactionError
	}
	defer func() {
		if p := recover(); p != nil {
			logrus.Errorf("事务回滚")
			tx.Rollback()
			// panic(p) // re-throw panic after Rollback
		} else if err != nil {
			logrus.Errorf("事务回滚")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			logrus.Info("事务提交")
		}
	}()

	// 1.添加数据到 vote_topic 表
	topicSql := `
insert into vote_topic(stu_id, title, description, deadline)  
values (?, ?, ?, ?)
`
	r, err := tx.Exec(topicSql, t.StuId, t.Title, t.Description, t.Deadline)
	if err != nil {
		logrus.Errorf("%s: %v\n", errno.MysqlInsertError, err)
		return errno.MysqlInsertError
	}
	TopicId, _ := r.LastInsertId()

	// 2.添加到 vote_set 表
	setSql := `
insert into vote_set(topic_id, select_type, anonymous, show_result, password) 
values (?, ?, ?, ?, ?)
`
	_, err = tx.Exec(setSql,
		TopicId, s.SelectType, s.Anonymous, s.ShowResult, s.Password)
	if err != nil {
		logrus.Errorf("%s: %v\n", errno.MysqlInsertError, err)
		return errno.MysqlInsertError
	}

	// 3.添加到 topic_option 表
	optionSql := `
insert into topic_option(topic_id, option_content) 
values (?, ?)
`
	for _, option := range o {
		_, err = tx.Exec(optionSql, TopicId, option.OptionContent)
		if err != nil {
			logrus.Errorf("%s: %v\n", errno.MysqlInsertError, err)
			return errno.MysqlInsertError
		}
	}

	// test rollback
	// panic("test rollback")
	return nil
}

func (d *TopicDao) QueryAllLimitWithTopicSet(page, size int) ([]*model.Topic, error) {
	if page <= 0 || size <= 0 {
		logrus.Errorf("%s, page: %d, size: %d\n",
			errno.MysqlLimitParamError.Error(), page, size)
		return nil, errno.MysqlLimitParamError
	}

	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select t.id,
       t.stu_id,
       t.title,
       t.description,
       t.deadline,
       t.review_status,
       t.create_time,
       vs.id,
       vs.topic_id,
       vs.select_type,
       vs.anonymous,
       vs.show_result,
       vs.password
from vote_topic t
         join vote_set vs on t.id = vs.topic_id
where t.delete_time is null
  /*and t.review_status = ?*/
limit ?,?
`

	limitOne := (page - 1) * size
	var ts []*model.Topic
	rows, err := mysql.Query(sql /*enum.TOPIC_REVIEW_PASS, */, limitOne, size)
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError.Error(), err.Error())
		return nil, errno.MysqlSelectError
	}
	for rows.Next() {
		var t model.Topic
		var s model.TopicSet
		if err := rows.Scan(
			&t.Id,
			&t.StuId,
			&t.Title,
			&t.Description,
			&t.Deadline,
			&t.ReviewStatus,
			&t.CreateTime,
			&s.Id,
			&s.TopicId,
			&s.SelectType,
			&s.Anonymous,
			&s.ShowResult,
			&s.Password); err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlScanError.Error(), err.Error())
			return nil, errno.MysqlScanError
		}
		t.TopicSets = &s
		ts = append(ts, &t)
	}
	//if err := mysql.Select(&ts, sql, enum.TOPIC_REVIEW_PASS, limitOne, size); err != nil {
	//	logrus.Errorf("%s: %s\n", errno.MysqlSelectError.Error(), err.Error())
	//	return nil, errno.MysqlSelectError
	//}

	return ts, nil
}

func (d *TopicDao) Count() (total int64, err error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return 0, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select count(*) from vote_topic 
where vote_topic.review_status = 1
`
	//var count int
	if err := mysql.Get(&total, sql); err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlSelectError, err)
		return 0, errno.MysqlSelectError
	}

	return total, nil
}

func (d *TopicDao) QueryByTitle(title string) ([]*model.Topic, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select t.id,
       t.stu_id,
       t.title,
       t.description,
       t.deadline,
       t.review_status,
       t.create_time,
       vs.id,
       vs.topic_id,
       vs.select_type,
       vs.anonymous,
       vs.show_result,
       vs.password
from vote_topic t
         join vote_set vs on t.id = vs.topic_id
where t.delete_time is null
  and t.review_status = ?
	and t.title = ?
`
	var ts []*model.Topic
	rows, err := mysql.Query(sql, enum.TOPIC_REVIEW_PASS, title)
	if err != nil {
		logrus.Warningf("%s: %s", errno.MysqlSelectNoData.Error(), err.Error())
		return nil, errno.MysqlSelectNoData
	}
	for rows.Next() {
		var s model.TopicSet
		var t model.Topic
		if err := rows.Scan(
			&t.Id,
			&t.StuId,
			&t.Title,
			&t.Description,
			&t.Deadline,
			&t.ReviewStatus,
			&t.CreateTime,
			&s.Id,
			&s.TopicId,
			&s.SelectType,
			&s.Anonymous,
			&s.ShowResult,
			&s.Password); err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlScanError.Error(), err.Error())
			return nil, errno.MysqlScanError
		}
		t.TopicSets = &s
		ts = append(ts, &t)
	}

	return ts, nil
}

func (d *TopicDao) QueryById(id string) (*model.Topic, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select t.id,
       t.stu_id,
       t.title,
       t.description,
       t.deadline,
       t.review_status,
       t.create_time,
       vs.id,
       vs.topic_id,
       vs.select_type,
       vs.anonymous,
       vs.show_result,
       vs.password
from vote_topic t
         join vote_set vs on t.id = vs.topic_id
where t.delete_time is null
  and t.review_status = ?
	and t.id = ?
`
	var t model.Topic
	rows, err := mysql.Query(sql, enum.TOPIC_REVIEW_PASS, id)
	if err != nil {
		logrus.Warningf("%s: %s", errno.MysqlSelectNoData.Error(), err.Error())
		return nil, errno.MysqlSelectNoData
	}
	for rows.Next() {
		var s model.TopicSet
		//var t model.Topic
		if err := rows.Scan(
			&t.Id,
			&t.StuId,
			&t.Title,
			&t.Description,
			&t.Deadline,
			&t.ReviewStatus,
			&t.CreateTime,
			&s.Id,
			&s.TopicId,
			&s.SelectType,
			&s.Anonymous,
			&s.ShowResult,
			&s.Password); err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlScanError.Error(), err.Error())
			return nil, errno.MysqlScanError
		}
		t.TopicSets = &s
	}

	return &t, nil
}

func (d *TopicDao) ShowResultById(id string) ([]*model.VoteResultVO, error) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		return nil, errno.MysqlConnectError
	}
	defer mysql.Close()

	sql := `
select id, option_content, number 
from topic_option
where topic_id = ?
`
	rows, err := mysql.Query(sql, id)
	if err != nil {
		logrus.Warningf("%s: %s", errno.MysqlSelectNoData.Error(), err.Error())
		return nil, errno.MysqlSelectNoData
	}

	var rs []*model.VoteResultVO
	for rows.Next() {
		var r model.VoteResultVO
		if err := rows.Scan(&r.OptionId, &r.OptionContent, &r.Votes); err != nil {
			logrus.Errorf("%s: %s\n", errno.MysqlScanError.Error(), err.Error())
			return nil, errno.MysqlScanError
		}
		rs = append(rs, &r)
	}
	return rs, nil
}
