package config

type Config struct {
	ProjectName string
	GoMod       string
}

func NewConfig(name, mod string) *Config {
	return &Config{
		ProjectName: name,
		GoMod:       mod,
	}
}
