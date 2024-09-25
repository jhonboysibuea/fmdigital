package config

import (
	"crud-app/logger"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

var config Config

type DbConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	Schema          string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime time.Duration
	Debug           bool
	EnableMigration bool
}

type JwtConfig struct {
	Ttl        int
	Key        string
	RefreshKey string
}
type Config struct {
	Db     *DbConfig
	Jwt    *JwtConfig
	Bucket *BucketConfig
	Redis  *RedisConfig
}
type BucketConfig struct {
	Endpoint string
	Space    string
	Key      string
	Secret   string
	Region   string
	Path     string
}

type RedisConfig struct {
	Host   string
	Pass   string
	DB     int
	Enable bool

	//key
}

func Db() *DbConfig {
	return config.Db
}

func Jwt() *JwtConfig {
	return config.Jwt
}
func Bucket() *BucketConfig {
	return config.Bucket
}

func Redis() *RedisConfig {
	return config.Redis
}
func LoadConfig() {
	configFileName := "config.json"
	if envProfile := os.Getenv("PROFILE"); len(envProfile) > 0 {
		logger.DefaultLogger().Info(fmt.Sprintf("started with profile %s", envProfile))
		configFileName = fmt.Sprintf("config.%s.json", envProfile)
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(configFileName)

	if err := viper.ReadInConfig(); err != nil {
		logger.DefaultLogger().Error(fmt.Sprintf("%s on reading default config from config.json", err.Error()))
	}

	config = Config{}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	if r, err := json.MarshalIndent(&config, "", "  "); err == nil {
		fmt.Println(string(r))
	}

}
