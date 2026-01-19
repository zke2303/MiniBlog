// Package main 应用的的入口
package main

import (
	"log"

	"mini-blog/internal/config"
)

// main 程序的主入口
func main() {
	// 1.读取配置文件
	_, err := config.ConfigurationInit()
	if err != nil {
		log.Fatal("error: %w", err)
	}
}
