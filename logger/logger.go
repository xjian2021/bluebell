package logger

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/xjian2021/bluebell/settings"
)

var logger *zap.Logger

func Init(c *settings.Logger) {
	//*************简单用法*************
	//{
	//	//logger, _ = zap.NewProduction() // json格式日志，常用于生产环境，方便解析
	//	logger, _ = zap.NewDevelopment() // 格式化日志，常用于开发环境，方便观察
	//	defer logger.Sync()              // 把缓存的日志写入磁盘
	//	logger.Info("logger init success...")
	//}

	//**************定制化用法**************
	{
		enc := getEncoder()
		ws := getLogWriter(c, c.Filepath+c.LogFilename)
		var l = new(zapcore.Level)
		err := l.UnmarshalText([]byte(c.LogLevel))
		if err != nil {
			panic(fmt.Errorf("l.UnmarshalText log_level fail err:%s", err.Error()))
		}
		core1 := zapcore.NewCore(enc, ws, l)
		// 把错误单独输出都一个日志文件中
		wsErr := getLogWriter(c, c.Filepath+c.ErrFilename)
		var el = new(zapcore.Level)
		err = el.UnmarshalText([]byte(c.ErrLevel))
		if err != nil {
			panic(fmt.Errorf("l.UnmarshalText err_level fail err:%s", err.Error()))
		}
		core2 := zapcore.NewCore(enc, wsErr, el)
		// 合并core
		core := zapcore.NewTee(core1, core2)
		// AddCaller 显示函数调用信息
		logger = zap.New(core, zap.AddCaller())
	}

	//****************自己配****************
	//{
	//	var err error
	//	config := zap.NewProductionConfig()
	//	level := zap.AtomicLevel{}
	//	err = level.UnmarshalText([]byte(c.LogLevel))
	//	if err != nil {
	//		panic(fmt.Errorf("level.UnmarshalText fail err:%s", err.Error()))
	//	}
	//	// ErrorOutputPaths是用来记录zap的内部错误，不是用于记录error级别的日志
	//	config.ErrorOutputPaths = []string{c.Filepath + c.ErrFilename}
	//	config.OutputPaths = []string{c.Filepath + c.LogFilename}
	//	config.EncoderConfig = getEncoderConfig()
	//	logger, err = config.Build()
	//	if err != nil {
	//		panic(fmt.Errorf("core.Build fail err:%s", err.Error()))
	//	}
	//}

	//sugarLogger = logger.Sugar() // printf式使用格式，使用方便但速度慢
	//defer sugarLogger.Sync()
	//sugarLogger.Info("sugarLogger init success...")
	logger.Info("logger init success...")
	// 把日志对象替换为zap包内的全局对象
	zap.ReplaceGlobals(logger)
}

func getEncoderConfig() zapcore.EncoderConfig {
	//cfg := zap.NewProductionEncoderConfig()
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("%s:%d", caller.Function, caller.Line))
	}
	return cfg
}

//getEncoder 获取日志编码器
func getEncoder() zapcore.Encoder {
	cfg := getEncoderConfig()
	//enc := zapcore.NewJSONEncoder(cfg)
	enc := zapcore.NewConsoleEncoder(cfg)
	return enc
}

//getLogWriter 获取日志写入对象
func getLogWriter(c *settings.Logger, filename string) zapcore.WriteSyncer {
	// 日志分割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    c.MaxSize,    // 单位M
		MaxAge:     c.MaxAge,     // 最大备份天数
		MaxBackups: c.MaxBackups, // 最大备份数量
		Compress:   c.Compress,   // 是否压缩
	}
	// 普通日志文件
	//file, err := os.OpenFile("./log/log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	//if err != nil {
	//	fmt.Println("os.OpenFile: ", err.Error())
	//	return nil
	//}
	// 输出到多个地方
	w := io.MultiWriter(lumberJackLogger, os.Stdout)
	ws := zapcore.AddSync(w)
	return ws
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
