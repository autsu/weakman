package test

import (
	"log"
	"testing"
	dao2 "vote/trash/internal/v1/dao"
)

func TestQueryAdminByUname(t *testing.T) {
	uname, err := dao2.AdminQueryByUname("root")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(uname)
}
