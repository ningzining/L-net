package bootstrap

type Option func(s *ServerBootstrap)

func WithServerIp(ip string) Option {
	return func(s *ServerBootstrap) {
		s.config.Host = ip
	}
}

func WithServerPort(port int) Option {
	return func(s *ServerBootstrap) {
		s.config.Port = port
	}
}
