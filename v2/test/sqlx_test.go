package test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
	"vote/v2/model"
)

func Test1(t *testing.T) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/vote?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	sql := "select * from vote_stu"
	var stu []model.Stu
	if err := db.Select(&stu, sql); err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("%+v\n", stu)
}
