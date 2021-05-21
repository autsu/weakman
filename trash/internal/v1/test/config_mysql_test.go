package test

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	config2 "vote/v1/config"
)

func TestMysqlConfig(t *testing.T) {
	dns, err := config2.NewMysqlConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(dns)
}
