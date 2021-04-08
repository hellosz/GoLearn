package parser

import (
	// "fmt"
	"hellosz.top/src/import/model"
	"log"
	"regexp"
	"time"
)

// 日志解析正则表达式
// var logExp = regexp.MustCompile(`^(\w+)\s(\S+)\s\[(.+?)\]\s(\S+)\s\-\s(\S+)\s(\S+)\s(\S+)\s"(.+?)"\s(\d+)[\s-]+(\d+)\s[\s\d]+"(.+?)"\s"(.+?)"[-\s]+(.+)$`)
var logExp = regexp.MustCompile(`^(\w+)\s(\S+)\s\[(.+?)\]\s(\S+)\s\S+\s(\S+)\s(\S+)\s(\S+)\s"(.+?)"\s(\d+)[\s-]+(\d+)\s[\s\d-]+"(.+?)"\s"(.+?)"[-\s]+(.+)`)

// Log 解析解析日志到model中，并返回结果
func Log(content []byte) model.Access {
	matches := logExp.FindSubmatch(content)

	// TODO 输出明细数据
	// for _, match := range matches {
	// 	fmt.Println(string(match))
	// }

	if len(matches) == 0 {
		log.Printf("日志解析失败, 日志原文：%s\n", string(content))

		return model.Access{}
	}

	// 解析时间字符串
	t, _ := time.Parse("02/Jan/2006:15:04:05 -0700", string(matches[3]))
	access := model.Access{
		BucketKey:  string(matches[1]),
		Bucket:     string(matches[2]),
		Datetime:   t.Format("2006-01-02 03:04:05"),
		IP:         string(matches[4]),
		Value1:     string(matches[5]),
		Method:     string(matches[6]),
		URI:        string(matches[7]),
		Request:    string(matches[8]),
		StatusCode: string(matches[9]),
		FileSize:   string(matches[10]),
		Value2:     string(matches[11]),
		Value3:     string(matches[12]),
		Value4:     string(matches[13]),
	}

	return access
}
