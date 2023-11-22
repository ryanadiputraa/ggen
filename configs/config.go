package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
	Mod  string
}

func LoadConfig(configType, filePath string) (*Config, error) {
	viper.SetConfigType(configType)
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}
