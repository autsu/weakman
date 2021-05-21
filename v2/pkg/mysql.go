package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"vote/v2/config"
)

func NewMysql() (*sqlx.DB, error) {
	dns, err := config.NewMysqlConfig()
	conn, err := sqlx.Connect("mysql", dns)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
