package settings

import "testing"

func TestSafeConfigRedactsPasswords(t *testing.T) {
	C = &conf{
		Name:      "web_app",
		Mode:      "dev",
		Version:   "v0.0.3",
		StartTime: "2022-09-26",
		Locale:    "zh",
		Port:      8081,
		MachineId: 1,
		Log: &Logger{
			LogLevel:    "debug",
			LogFilename: "log.log",
			ErrLevel:    "error",
			ErrFilename: "err.log",
			Filepath:    "./log/",
			MaxSize:     200,
			MaxAge:      30,
			MaxBackups:  7,
			Compress:    false,
		},
		Mysql: &Mysql{
			Host:         "127.0.0.1",
			Port:         3306,
			User:         "root",
			Password:     "secret",
			Dbname:       "db1",
			MaxOpenConns: 100,
			MaxIdleConns: 10,
		},
		Redis: &Redis{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "redis-secret",
			Db:       0,
			PoolSize: 100,
		},
	}

	got := SafeConfig()
	mysqlConfig := got["mysql"].(map[string]interface{})
	redisConfig := got["redis"].(map[string]interface{})

	if mysqlConfig["password"] == "secret" {
		t.Fatal("mysql password was not redacted")
	}
	if redisConfig["password"] == "redis-secret" {
		t.Fatal("redis password was not redacted")
	}
	if mysqlConfig["password"] != "******" {
		t.Fatalf("unexpected mysql password value: %v", mysqlConfig["password"])
	}
	if redisConfig["password"] != "******" {
		t.Fatalf("unexpected redis password value: %v", redisConfig["password"])
	}
}
