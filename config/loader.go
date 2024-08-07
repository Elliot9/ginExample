package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("config 加載 .env 時發生錯誤: %v\n", err)
	}

	AppSetting = &App{
		Url: loadEnv("APP_URL", "localhost:8080"),
		Env: loadEnv("APP_ENV", "debug"),
	}

	WDbSetting = &Database{
		Host:     loadEnv("DB_HOST", "127.0.0.1"),
		Port:     loadEnv("DB_PORT", 3306),
		Name:     loadEnv("DB_DATABASE", "ginExample"),
		UserName: loadEnv("DB_USERNAME", "root"),
		Password: loadEnv("DB_PASSWORD", ""),
	}

	RDbSetting = &Database{
		Host:     loadEnv("DB_READONLY_HOST", WDbSetting.Host),
		Port:     loadEnv("DB_READONLY_PORT", WDbSetting.Port),
		Name:     loadEnv("DB_READONLY_DATABASE", WDbSetting.Name),
		UserName: loadEnv("DB_READONLY_USERNAME", WDbSetting.UserName),
		Password: loadEnv("DB_READONLY_PASSWORD", WDbSetting.Password),
	}

	RedisSetting = &Redis{
		Host:     loadEnv("REDIS_HOST", "127.0.0.1"),
		Port:     loadEnv("REDIS_PORT", 6379),
		Password: loadEnv("REDIS_PASSWORD", ""),
	}
}

func loadEnv[T ~string | ~int | ~float64 | ~bool](key string, defaultValue T) T {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defaultValue
	}

	var result any
	var err error

	switch any(defaultValue).(type) {
	case int:
		var parsedValue int
		parsedValue, err = strconv.Atoi(envValue)
		result = parsedValue
	case float64:
		var parsedValue float64
		parsedValue, err = strconv.ParseFloat(envValue, 64)
		result = parsedValue
	case bool:
		var parsedValue bool
		parsedValue, err = strconv.ParseBool(envValue)
		result = parsedValue
	case string:
		result = envValue
	default:
		err = fmt.Errorf("不支持的類型")
	}

	if err != nil {
		fmt.Printf("加載 %s 時發生錯誤: %v，使用預設值 %v\n", key, err, defaultValue)
		return defaultValue
	}

	return result.(T)
}
