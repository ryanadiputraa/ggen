package config

type Config struct {
	ProjectName string
	GoMod       string
	GgenDir     string
	ProjectDir  string
}

func NewConfig(name, mod string) *Config {
	return &Config{
		ProjectName: name,
		GoMod:       mod,
	}
}