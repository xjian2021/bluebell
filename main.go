package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	validatorPkg "bluebell/pkg/validator-trans"
	"bluebell/routes"
	"bluebell/settings"
)

func main() {
	// 加载配置
	settings.Init()
	// 初始化日志
	logger.Init(settings.C.Log)
	defer zap.L().Sync()
	// 初始化mysql
	mysql.Init(settings.C.Mysql)
	defer mysql.Close()
	// 初始化redis
	redis.Init(settings.C.Redis)
	defer redis.Close()
	// 初始化分布式id
	if err := snowflake.Init(settings.C.StartTime, int64(settings.C.MachineId)); err != nil {
		zap.L().Fatal("snowflake.Init", zap.Error(err))
	}

	// 初始化validator全局翻译器
	if err := validatorPkg.InitTrans(settings.C.Locale); err != nil {
		zap.L().Fatal("validator-trans.InitTrans", zap.Error(err))
	}

	// 注册路由
	r := routes.Setup()

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.C.Port),
		Handler: r,
	}
	go func() {
		if err := svr.ListenAndServe(); err != nil {
			zap.L().Info("svr.ListenAndServe", zap.Error(err))
		}
	}()

	zap.L().Info("hello world!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	zap.L().Debug("close server...", zap.String("signal", s.String()))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		zap.L().Fatal("svr.Shutdown", zap.Error(err))
	}
	zap.L().Info("server exit!")
}
