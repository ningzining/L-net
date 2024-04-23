package conf

type ServerOption func(s *ServerConfig)

func WithServerPort(port int) ServerOption {
	return func(s *ServerConfig) {
		s.Port = port
	}
}
