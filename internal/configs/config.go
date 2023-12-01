package configs

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/writer"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/configs", name)
	if err := writer.CreateDirectory(path); err != nil {
		return err
	}

	if err := writer.WriteToFile(ymlConfigTemplate(), path, "config.yml"); err != nil {
		return err
	}

	return writer.WriteToFile(template(mod), path, "config.go")
}

func template(mod string) string {
	return `package configs

import (
    "github.com/spf13/viper"
)

type Config struct {
    *Server
}

type Server struct {
    Port string ` + "`mapstructure:\"port\"`" + `
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
}`
}

func ymlConfigTemplate() string {
	return `server:
    port: :8080`
}
