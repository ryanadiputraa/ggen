package config

type Config struct {
	ProjectName string
	GoMod       string
	OriginPath  string
}

func NewConfig(name, mod, originPath string) *Config {
	return &Config{
		ProjectName: name,
		GoMod:       mod,
		OriginPath:  originPath,
	}
}
