package config

import (
	"time"

	"github.com/linkeunid/api.linkeun.com/pkg/env"

	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Port       int64
	Env        string
	AppSalt    string
	JWTSecret  string
	JWTExpires int64

	Dsn       string
	SentryDsn string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	cfg.Env = env.GetString("ENV", "production")

	if cfg.Env == "development" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	cfg.Host = env.GetString("HOST", "localhost")
	cfg.Port = env.GetInt("PORT", 4444)
	cfg.AppSalt = env.GetString("APP_SALT", "salt")
	cfg.JWTSecret = env.GetString("JWT_SECRET", "secret")
	cfg.JWTExpires = int64((time.Duration(env.GetInt("JWT_EXPIRES", 1)) * time.Hour).Seconds())

	cfg.Dsn = env.GetString("DSN", "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	cfg.SentryDsn = env.GetString("SENTRY_DSN", "")

	return cfg, nil
}
