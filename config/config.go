package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Server struct {
	Port int
}

type Db struct {
	Port       int
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbSslmode  string
}

type Config struct {
	Server *Server
	Db     *Db
}

var once sync.Once
var configInstance *Config

func Setup() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&configInstance); err != nil {
		panic(err)
	}
}
