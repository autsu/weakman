package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	//"github.com/sirupsen/logrus"
	"vote/config"
	//"vote/errno"
)

var (
	mysqlConn *sqlx.DB
	errs      error
)

func init() {
	conn, err := newMysqlConn()
	if err != nil {
		// logrus.Errorf("%s: %s\n", errno.MysqlConnectError, err)
		// panic("connect mysql error: " + err.Error())
		errs = err
	}
	mysqlConn = conn
}

func newMysqlConn() (*sqlx.DB, error) {
	dns, err := config.NewMysqlConfig()
	if err != nil {
		return nil, err
	}

	conn, err := sqlx.Connect("mysql", dns)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(16)
	conn.SetMaxIdleConns(8)
	return conn, nil
}

func GetConn() (*sqlx.DB, error) {
	return mysqlConn, errs
}
