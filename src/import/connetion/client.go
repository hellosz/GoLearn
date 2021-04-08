package connetion

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// Connection 数据库连接
var Connection *sql.DB

// Get 创建数据库连接
func Get() *sql.DB {
	Connection, err := sql.Open("mysql", "root:patpat@tcp(192.168.11.131)/mms")
	if err != nil {
		log.Printf("数据库连接失败，原因:%v", err)
	}

	// 配置连接属性
	Connection.SetConnMaxIdleTime(time.Minute * 3)
	Connection.SetMaxOpenConns(10)
	Connection.SetMaxIdleConns(10)

	return Connection
}

// Close 关闭连接
func Close() {
	if Connection != nil {
		Connection.Close()
	}
}
