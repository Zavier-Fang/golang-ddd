package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath("./server/conf")
	Config.AddConfigPath(".")
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig(); err != nil{
		log.Fatal("load config failed! err: ", err.Error())
	}
}
