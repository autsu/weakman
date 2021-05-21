package service

import (
	"log"
	dao2 "vote/v1/dao"
	errno2 "vote/v1/errno"
	model2 "vote/v1/model"
)

func CollegeGetAll() ([]*model2.College, error) {
	college, err := dao2.CollegeQueryAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return college, nil
}

func CollegeAdd(c *model2.College) error {
	exist := CollegeIsExist(c.Name)
	if exist {
		return errno2.DB_COLLEGE_NAME_EXISIT
	}

	n, err := dao2.CollegeAdd(c.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	if n < 1 {
		return errno2.DB_INSERT_ERROR
	}

	return nil
}

func CollegeSoftDelete(id string) (n int64, err error) {
	n, err = dao2.CollegeSoftDelete(id)
	if err != nil {
		return 0, err
	}
	return
}

func CollegeUpdate(coll *model2.College) (n int64, err error) {
	n, err = dao2.CollegeUpdate(coll)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func CollegeQueryByIdWithProfession(id string) (*model2.College, error) {
	p, err := dao2.CollegeQueryByIdWithProfession(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return p, nil
}

func CollegeIsExist(name string) bool {
	cc, err := dao2.CollegeQueryByName(name)
	if err != nil { // the result set is empty
		log.Println(err)
		return false
	}

	if cc != nil {
		return true
	}
	return false
}

