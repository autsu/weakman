package test

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	pkg2 "vote/v1/pkg"
)

func TestMysqlPing(t *testing.T) {
	mysql, err := pkg2.NewMysql()
	if err != nil {
		log.Fatalln(err)
	}
	if err := mysql.Ping(); err != nil {
		log.Fatalln(err)
	}
}
