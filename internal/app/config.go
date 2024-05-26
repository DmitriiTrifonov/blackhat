package app

var DefaultConfig = defaultConfig()

type Config struct {
}

func defaultConfig() *Config {
	return &Config{}
}
