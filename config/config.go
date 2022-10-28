package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/vodeacloud/hr-api/pkg/logger"
	"time"
)

type Config struct {
	OAuthSecret *OAuthSecret

	AppEnv      string `envconfig:"APP_ENV" default:"local"`
	AppTimezone string `envconfig:"APP_TIMEZONE" default:"Asia/Jakarta"`
	AppPort     string `envconfig:"APP_PORT" defaut:"80"`

	DBHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort     string `envconfig:"DB_PORT" default:"3306"`
	DBDatabase string `envconfig:"DB_DATABASE" required:"true"`
	DBUsername string `envconfig:"DB_USERNAME" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`

	RedisHost     string `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	RedisPort     string `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`

	UserPasswordMinLength int `envconfig:"USER_PASSWORD_MIN_LENGTH" default:"8"`
}

var cfg *Config

func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	cfg = &Config{}
	envconfig.MustProcess("", cfg)
	cfg.OAuthSecret.validateKey()
	return cfg
}

func GetLocation() *time.Location {
	loc, err := time.LoadLocation(GetConfig().AppTimezone)
	if err != nil {
		logger.Printf("failed to get location: %v", err)
	}
	return loc
}

func IsProd() bool {
	return GetConfig().AppEnv == "production"
}
