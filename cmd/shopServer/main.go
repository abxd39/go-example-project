package main

import (
	"context"
	"example-project/internal/api"
	"example-project/internal/domain"
	"example-project/internal/repository"
	"example-project/internal/service"
	"example-project/pkg/config"
	"example-project/pkg/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 创建仓库
	reminderRepo := repository.NewReminderRepository(db)

	// 创建服务
	reminderService := service.NewReminderService(reminderRepo)

	// 创建API处理器
	reminderHandler := api.NewReminderHandler(reminderService)

	// 设置路由
	router := setupRouter(cfg)
	api.RegisterRoutes(router, reminderHandler)

	srv := &http.Server{
		Addr:    cfg.Server.Address,
		Handler: router,
	}

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func initDatabase(cfg domain.Config) (*gorm.DB, error) {
	// 实际的数据库初始化代码
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 配置连接池
	db.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Hour)

	// 自动迁移表结构
	if err := db.AutoMigrate(&domain.Reminder{}).Error; err != nil {
		return nil, err
	}

	return db, nil
}

func setupRouter(cfg domain.Config) *gin.Engine {
	if cfg.Server.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// 添加中间件
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.RequestID())
	router.Use(middleware.RateLimiter(cfg.RateLimit))
	registerHealthRoutes(router)
	return router
}

func registerHealthRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"version":   "1.0.0",
			"timestamp": time.Now().Unix(),
		})
	})
}
