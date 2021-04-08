package connetion

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"hellosz.top/src/import/utils"
)

// Connection 数据库连接
var Connection *sql.DB

// Get 创建数据库连接
func Get(config utils.Connection) *sql.DB {
	if Connection != nil {
		return Connection
	}

	// 根据参数配置数据数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	Connection, err := sql.Open("mysql", dsn)
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
