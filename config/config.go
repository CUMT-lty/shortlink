package config

import (
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

var Conf *Config

type Config struct {
	Common Common
}

func init() {
	InitConfig()
}

func InitConfig() {
	once.Do(func() {
		configFilePath := "./config" // TODO: 这个路径该怎么搞
		viper.SetConfigType("toml")
		viper.SetConfigName("/common")
		viper.AddConfigPath(configFilePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		Conf = new(Config)
		err = viper.Unmarshal(&Conf.Common)
		if err != nil {
			panic(err)
		}
	})
}
