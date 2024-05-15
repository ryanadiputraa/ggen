package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"port"`
	DBHost     string `mapstructure:"db_host"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
	DBPort     string `mapstructure:"db_port"`
}

func NewConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config/")

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.New("config file not found")
			return
		}
		return
	}

	err = viper.Unmarshal(&config)
	return
}
