package config

import (
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/log"

	"github.com/spf13/viper"
)

var AppConfig *appConfig

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../.config/")
	viper.AddConfigPath("/var/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.Error("Config did not read")
		return
	}
	var cfg appConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Logger.Error("Config did not read")
		return
	}
	AppConfig = &cfg
}
