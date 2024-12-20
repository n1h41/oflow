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

type AWS struct {
	ClientId       string
	ClientSecret   string
	IdentityPoolId string
	AccountId      string
}

type Config struct {
	Server *Server
	Db     *Db
	AWS    *AWS
}

var (
	once           sync.Once
	configInstance *Config
)

func Setup() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config/")
		viper.AddConfigPath("../../config/")
		viper.AutomaticEnv()
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&configInstance); err != nil {
		panic(err)
	}
	return configInstance
}
