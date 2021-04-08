package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config 主配置文件
type Config struct {
	LogDir      string     `yaml:"logDir"`
	WorkerCount int        `yaml:"workerCount"`
	Connection  Connection `yaml:"connection"`
}

// Connection 数据库连接配置
type Connection struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

// Get 读取配置文件，返回Config对象
func (c *Config) Get(filepath string) *Config {
	// 读取配置文件
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("读取配置文件(%s)失败, 原因：%v", filepath, err)
	}

	// 解析配置
	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatalf("解析yml文件失败，原因：%v", err)
	}

	// 返回结果
	return c
}
