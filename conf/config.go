package conf

type Config struct {
	Host           string
	Port           int
	MaxPackageSize int
	MaxConnSize    int
}

func DefaultConfig() *Config {
	return &Config{
		Host:           "0.0.0.0",
		Port:           8888,
		MaxPackageSize: 1024,
		MaxConnSize:    1024 * 10,
	}
}
