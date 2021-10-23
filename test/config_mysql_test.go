package test

import (
	"log"
	"testing"
	"vote/config"

	_ "github.com/go-sql-driver/mysql"
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
