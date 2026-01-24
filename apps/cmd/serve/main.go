// Package main 应用的的入口
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"mini-blog/internal/config"
	"mini-blog/internal/controller"
	"mini-blog/internal/middleware"
	"mini-blog/internal/pkg/common"
	"mini-blog/internal/repository"
	"mini-blog/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// main 程序的主入口
func main() {
	// 1.读取配置文件
	cfg, err := config.ConfigurationInit()
	if err != nil {
		log.Fatal("读取配置文件失败: %w", err)
	}

	// 注册 全局翻译器
	if err := common.InitTranslation(); err != nil {
		log.Fatal("初始化翻译器错误: ", err)
	}

	// 2.连接数据库
	newLogger := logger.New(
		log.New(os.Stdout, "\n\r", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)
	db, err := gorm.Open(postgres.Open(cfg.Datasource.Postgres.Dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("数据库连接错误: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Datasource.Redis.Addr,
		DB:   cfg.Datasource.Redis.DB,
	})
	// 3.创建 repository 实例对象
	userRepo := repository.NewUserRepository()
	blogRepo := repository.NewBlogRepository(rdb)
	// 4.创建 service 实例对象
	userService := service.NewUserService(db, userRepo)
	blogService := service.NewBlogService(blogRepo, db)
	// 5.创建 controller 实例对象
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)
	blogController := controller.NewBlogController(blogService)
	// 6.配置路由
	r := setupRouter(userController, authController, blogController)

	// 7.设置 gin 启动模式
	gin.SetMode(cfg.Serve.Mode)

	// 8.启动服务
	addr := fmt.Sprintf("%s:%d", "localhost", cfg.Serve.Port)
	log.Fatal(r.Run(addr))
}

// setupRouter 设置路由
func setupRouter(
	userController *controller.UserController,
	authController *controller.AuthController,
	blogController *controller.BlogController,
) *gin.Engine {
	// 1.创建 gin 示例对象
	r := gin.Default()

	// 2.配置路由
	v1 := r.Group("/api/v1", middleware.ErrorHandlerMiddleware(common.GlobalTrans))
	// 3.分组路由
	{
		// 免登入
		auth := v1.Group("")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)

			blogs := v1.Group("/blogs")
			{
				blogs.GET("/", blogController.ListBlogs)
				blogs.GET("/:id", blogController.GetBlogDetail)
			}
		}

		// 需要登入
		protected := v1.Group("", middleware.AuthMiddleware())
		{
			users := protected.Group("/users")
			{
				users.GET("/profile", userController.Profile)
				users.PUT("/update", userController.Update)
			}

			blogs := protected.Group("/blogs")
			{
				blogs.POST("/", blogController.Create)
				blogs.DELETE("/:id", blogController.Delete)
			}
		}
	}

	return r
}
