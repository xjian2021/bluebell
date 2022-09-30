package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/settings"
)

var db *sqlx.DB

func Init(c *settings.Mysql) {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Dbname,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("sqlx.Connect fail err:%s", err.Error()))
	}
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)
	if err = db.Ping(); err != nil {
		panic(fmt.Errorf("db.Ping fail err:%s", err.Error()))
	}
	zap.L().Info("init mysql success...")
}

func Close() {
	_ = db.Close()
}
