package service

import (
	"strconv"
	"vote/v2/dao"
	"vote/v2/errno"
	"vote/v2/model"
	"vote/v2/pkg"
	"vote/v2/tool"
)

func StuLogin(username, password string) (token string, err error) {
	stu, err := dao.StuQueryByUsername(username)
	if err != nil {
		return "", err
	}

	if !tool.MD5Equal(password, stu.Password) {
		return "", errno.LoginPasswordWrong
	}
	jwt, err := pkg.CreateJwt(strconv.Itoa(stu.Id))
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func StuRegister(s *model.Stu) (int64, error) {
	if StuIsExist(s) {
		return 0, errno.RegisterPhoneIsExist
	}
	encryPwd, err := tool.NewMD5(s.Password)
	if err != nil {
		return 0, err
	}

	s.Password = encryPwd	// 加密后的密码
	n, err := dao.StuInsert(s)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func StuIsExist(s *model.Stu) bool {
	_, err := dao.StuQueryByPhone(s.Phone)
	if err != nil {
		return false
	}
	return true
}

func StuGetInfoByToken(token string) (*model.Stu, error) {
	t, err := pkg.ParseTokenWithBearer(token)
	if err != nil {
		return nil, err
	}
	stu, err := dao.StuQueryById(t.Id)
	if err != nil {
		return nil, err
	}
	return stu, nil
}
