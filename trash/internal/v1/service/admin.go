package service

import (
	"log"
	dao2 "vote/trash/internal/v1/dao"
	errno2 "vote/trash/internal/v1/errno"
	pkg2 "vote/trash/internal/v1/pkg"
	tool2 "vote/trash/internal/v1/tool"
)

func AdminLogin(username, password string) (token string, err error) {
	admin, err := dao2.AdminQueryByUname(username)
	if err != nil {
		log.Println(err)
		return "", err
	}
	equal := tool2.MD5Equal(password, admin.Password)
	if equal {
		jwt := pkg2.CreateJwt(admin.Id)
		return jwt, nil
	}
	return "", errno2.PASSWORD_NOT_EQUAL
}
