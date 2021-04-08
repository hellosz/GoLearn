package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Access struct {
	BucketKey  string
	Bucket     string
	Datetime   string
	IP         string
	Value1     string
	Method     string
	URI        string
	Request    string
	StatusCode string
	FileSize   string
	Value2     string
	Value3     string
	Value4     string
}

// Create 保存数据
func (a *Access) Create(db *sql.DB) (int, error) {
	result, err := db.Exec("INSERT INTO `mms_aws_access_log` (`bucket_key`, `bucket`, `datetime`, `ip`, `uri`, `request` "+
		", `status_code`, `file_size`, `value1`, `value2`, `value3`, `value4`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", a.BucketKey, a.Bucket, a.Datetime,
		a.IP, a.URI, a.Request, a.StatusCode, a.FileSize, a.Value1, a.Value2, a.Value3, a.Value4)
	if err != nil {
		log.Printf("发生错误%v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("发生错误%v", err)
		return 0, err
	}
	return int(id), nil
}

// Query 查询数据
func (a *Access) Query(id int, db *sql.DB) Access {
	err := db.QueryRow("select username from mms_sso_users where id = ?", id).Scan(&a.Bucket)
	if err != nil {
		log.Printf("发生错误%v", err)
	}

	return *a
}
