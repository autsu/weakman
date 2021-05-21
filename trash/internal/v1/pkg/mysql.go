package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	config2 "vote/trash/internal/v1/config"
)

func NewMysql() (*sqlx.DB, error) {
	dns, err := config2.NewMysqlConfig()
	conn, err := sqlx.Connect("mysql", dns)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
