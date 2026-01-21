// Package config 使用viper读取配置文件
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration 全局配置类
type Configuration struct {
	Serve      Serve      `mapstructure:"serve" json:"serve"`
	Datasource Datasource `mapstructure:"datasource" json:"datasource"`
}

// Serve Gin服务配置
type Serve struct {
	Port int    `mapstructure:"port" json:"port" xml:"port"`
	Mode string `mapstructure:"mode" json:"mode" xml:"mode"`
}

// Datasource 数据源配置
type Datasource struct {
	Postgres Postgres `mapstructure:"postgres" json:"postgres"`
	Redis    Redis    `mapstructure:"redis" json:"redis"`
}

// Postgres Postgres数据库配置
type Postgres struct {
	Dsn string `mapstructure:"dsn" json:"dsn"`
}

// Redis Redis 配置
type Redis struct {
	Addr string `mapstructure:"addr" json:"addr"`
	DB   int    `mapstructure:"db" json:"db"`
}

// ConfigurationInit 初始化Configuration对象
func ConfigurationInit() (*Configuration, error) {
	// 创建 viper 实例对象
	v := viper.New()
	// 2.配置 viper
	v.SetConfigType("yml")
	v.SetConfigName("config")
	v.AddConfigPath("./config")
	// 3.尝试读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error reading config file: %w", err)
	}
	// 4.读取配置文件成功，映射到 Configuration类中
	var cfg Configuration
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	// 5.返回
	return &cfg, nil
}
