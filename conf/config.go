package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/clippingkk/common/settings"
	"github.com/go-sql-driver/mysql"
)

var (
	appConfig *AppConfig

	// default params
	appEnv   = "dev"
	fileName = "app.dev.json"
)

func init() {
	if env := os.Getenv("APP_ENV"); env != "" {
		appEnv = strings.ToLower(env)
		fileName = fmt.Sprintf("app.%s.json", appEnv)
	}

	UniversalConfig()
}

// AppConfig ...
type AppConfig struct {
	AppName  string         `json:"app_name"`
	AppEnv   string         `json:"app_env"`
	Debug    bool           `json:"debug"`
	Database DatabaseConfig `json:"database"`
	Redis    RedisConfig    `json:"redis"`
}

// DatabaseConfig ....
type DatabaseConfig struct {
	DSN string `json:"dsn"`
	Cfg *mysql.Config
}

// RedisConfig ...
type RedisConfig struct {
	Addr       string `json:"addr" env:"REDIS_ADDR"`
	DB         int    `json:"db" env:"REDIS_DB"`
	MaxRetries int    `json:"max_retries"`
}

// UniversalConfig ...
func UniversalConfig() {
	cfg := &AppConfig{}
	if err := settings.Initialize(fileName, cfg); err != nil {
		panic(err)
	}

	appConfig = cfg
}

// GetConfig ...
func GetConfig() *AppConfig {
	return appConfig
}