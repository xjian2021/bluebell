package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var C = new(conf)

type (
	conf struct {
		Name      string  `mapstructure:"name"`
		Mode      string  `mapstructure:"mode"`
		Version   string  `mapstructure:"version"`
		StartTime string  `mapstructure:"start_time"`
		Locale    string  `mapstructure:"locale"`
		Port      int     `mapstructure:"port"`
		MachineId int     `mapstructure:"machine_id"`
		Log       *Logger `mapstructure:"log"`
		Mysql     *Mysql  `mapstructure:"mysql"`
		Redis     *Redis  `mapstructure:"redis"`
	}
	Logger struct {
		LogLevel    string `mapstructure:"log_level"`
		LogFilename string `mapstructure:"log_filename"`
		ErrLevel    string `mapstructure:"err_level"`
		ErrFilename string `mapstructure:"err_filename"`
		Filepath    string `mapstructure:"filepath"`
		MaxSize     int    `mapstructure:"max_size"`
		MaxAge      int    `mapstructure:"max_age"`
		MaxBackups  int    `mapstructure:"max_backups"`
		Compress    bool   `mapstructure:"compress"`
	}
	Mysql struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		Dbname       string `mapstructure:"db_name"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
	}
	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
		PoolSize int    `mapstructure:"pool_size"`
	}
)

func Init() {
	//viper.SetConfigFile("./conf/config.yaml") // 指定顶配置文件，不受AddConfigPath影响，需要指定路径
	viper.SetConfigName("config")  // 指定配置文件名称(不需要带后缀)
	viper.AddConfigPath("./conf/") // 添加查找配置文件的路径
	// 根据添加配置文件路径的顺序来作为查找优先级排序
	//viper.AddConfigPath("./")
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		panic(fmt.Errorf("viper.ReadInConfig fail err:%s", err.Error()))
	}
	err = viper.Unmarshal(C)
	if err != nil {
		panic(fmt.Errorf("viper.Unmarshal fail err:%s", err.Error()))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("config change!!!", zap.String("name", in.Name), zap.String(
			"op",
			in.Op.String(),
		))
		err = viper.Unmarshal(C)
		if err != nil {
			zap.L().Error(
				"viper.Unmarshal fail",
				zap.String("err", err.Error()),
			)
		}
	})
}
