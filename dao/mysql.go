package dao

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var MysqlEngine *Engine

func init() {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "123456789",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "go_blog",
		ParseTime:            true,
		AllowNativePasswords: true,
		Loc:                  time.Local,
	}

	MysqlEngine, _ = NewEngine("mysql", cfg.FormatDSN())
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	MysqlEngine.db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	MysqlEngine.db.SetMaxOpenConns(100)
	// 连接最大存活时间
	MysqlEngine.db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	MysqlEngine.db.SetConnMaxIdleTime(time.Minute * 1)
}
