package test

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"vote/pkg/mysql"
)

func TestMysqlPing(t *testing.T) {
	db, err := mysql.GetConn()
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
}
