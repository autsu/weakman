package pkg

import (
	"vote/v2/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMysql() (*sqlx.DB, error) {
	dns, err := config.NewMysqlConfig()
	conn, err := sqlx.Connect("mysql", dns)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(16)
	conn.SetMaxIdleConns(8)
	return conn, nil
}
