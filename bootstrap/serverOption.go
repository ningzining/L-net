package bootstrap

type Option func(s *Server)

func WithServerIp(ip string) Option {
	return func(s *Server) {
		s.config.Host = ip
	}
}

func WithServerPort(port int) Option {
	return func(s *Server) {
		s.config.Port = port
	}
}
