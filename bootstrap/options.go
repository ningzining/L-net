package bootstrap

type Option func(s *ServerBootstrap)

func WithIp(ip string) Option {
	return func(s *ServerBootstrap) {
		s.ip = ip
	}
}

func WithPort(port int) Option {
	return func(s *ServerBootstrap) {
		s.port = port
	}
}
