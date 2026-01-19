// Package main 应用的的入口
package main

import (
	"log"

	"mini-blog/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// main 程序的主入口
func main() {
	// 1.读取配置文件
	cfg, err := config.ConfigurationInit()
	if err != nil {
		log.Fatal("error: %w", err)
	}
	// 2.连接数据库
	_, err = gorm.Open(postgres.Open(cfg.Datasource.Postgres.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接错误: %w", err)
	}
}
