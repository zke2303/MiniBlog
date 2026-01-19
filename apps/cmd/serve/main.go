// Package main 应用的的入口
package main

import (
	"fmt"
	"log"

	"mini-blog/internal/config"
	"mini-blog/internal/controller"
	"mini-blog/internal/repository"
	"mini-blog/internal/service"

	"github.com/gin-gonic/gin"
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
	db, err := gorm.Open(postgres.Open(cfg.Datasource.Postgres.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接错误: %w", err)
	}
	// 3.创建 repository 实例对象
	userRepo := repository.NewUserRepository()

	// 4.创建 service 实例对象
	userService := service.NewUserService(db, userRepo)

	// 5.创建 controller 实例对象
	userController := controller.NewUserController(userService)

	// 6.配置路由
	r := setupRouter(*userController)

	// 7.设置 gin 启动模式
	gin.SetMode(cfg.Serve.Mode)

	// 8.启动服务
	addr := fmt.Sprintf("%s:%d", "localhost", cfg.Serve.Port)
	log.Fatal(r.Run(addr))
}

// setupRouter 设置路由
func setupRouter(
	userController controller.UserController,
) *gin.Engine {
	// 1.创建 gin 示例对象
	r := gin.Default()

	// 2.配置路由
	v1 := r.Group("/api/v1")
	// 3.分组路由
	{
		public := v1.Group("/public")
		{
			users := public.Group("/users")
			{
				users.POST("", userController.Create)
			}
		}
	}

	return r
}
