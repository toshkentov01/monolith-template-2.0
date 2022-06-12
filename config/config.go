package config

import (
	"os"
	"sync"
	"time"

	"github.com/spf13/cast"
)

var (
	instance *Configuration
	once     sync.Once
)

// Configuration ...
type Configuration struct {
	Environment                string
	PostgresHost               string
	PostgresPort               int
	PostgresUser               string
	PostgresPassword           string
	PostgresDB                 string
	ServerPort                 int
	ServerHost                 string
	LogLevel                   string
	ServiceDir                 string
	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration
	RedisHost                  string
	RedisPort                  int
	CasbinConfigPath           string
	MiddleWareRolesPath        string
	CtxTimeout                 int
	SigninKey                  string
	ServerReadTimeout          int
	JWTSecretKey               string
	JWTSecretKeyExpireMinutes  int
	JWTRefreshKey              string
	JWTRefreshKeyExpireHours   int
}

// load ...
func load() *Configuration {
	return &Configuration{
		ServerHost:          cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:          cast.ToInt(getOrReturnDefault("SERVER_PORT", "8000")),
		Environment:         cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		PostgresHost:        cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:        cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDB:          cast.ToString(getOrReturnDefault("POSTGRES_DB", "")),
		PostgresUser:        cast.ToString(getOrReturnDefault("POSTGRES_USER", "")),
		PostgresPassword:    cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
		LogLevel:            cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug")),
		ServiceDir:          cast.ToString(getOrReturnDefault("CURRENT_DIR", "")),
		CtxTimeout:          cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7)),
		RedisHost:           cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost")),
		RedisPort:           cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379)),
		CasbinConfigPath:    cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./config/rbac_model.conf")),
		MiddleWareRolesPath: cast.ToString(getOrReturnDefault("MIDLEWARE_ROLES_PATH", "./config/models.csv")),
		SigninKey:           cast.ToString(getOrReturnDefault("SIGNIN_KEY", "")),
		ServerReadTimeout:   cast.ToInt(getOrReturnDefault("SERVER_READ_TIMEOUT", "")),

		JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
		JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
		JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
		JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

//Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}
