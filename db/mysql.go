package db

import (
	"com.jxtech.gather/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var _conn sql.DB

func init() {
	fmt.Println("初始化Mysql连接池：")
	param := config.Get()
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		param.Mysql.User, param.Mysql.Pass, "tcp", param.Mysql.Host, param.Mysql.Port, param.Mysql.Db)
	DB, err := sql.Open("db", dsn)
	if err != nil {
		fmt.Printf("Open db failed,gerr:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(20)                   //设置闲置连接数

	_conn = *DB
}

func Conn() *sql.DB {
	return &_conn
}
