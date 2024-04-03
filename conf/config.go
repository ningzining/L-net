package conf

type Config struct {
	Host string
	Port int
}

func DefaultConfig() *Config {
	return &Config{
		Host: "0.0.0.0",
		Port: 8888,
	}
}
