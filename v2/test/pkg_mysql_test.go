package test

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"vote/v2/pkg"
)

func TestMysqlPing(t *testing.T) {
	mysql, err := pkg.NewMysql()
	if err != nil {
		log.Fatalln(err)
	}
	if err := mysql.Ping(); err != nil {
		log.Fatalln(err)
	}
}
