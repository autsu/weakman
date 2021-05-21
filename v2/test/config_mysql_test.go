package test

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"vote/v2/config"
)

func init() {
	log.SetFlags(log.Llongfile | log.Ldate)
}

func TestMysqlConfig(t *testing.T) {
	dns, err := config.NewMysqlConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(dns)
}
