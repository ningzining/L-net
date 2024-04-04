package bootstrap

import (
	"errors"
	"net"

	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/iface"
)

type ServerBootstrap struct {
	ip      string
	port    int
	decoder decoder.Decoder
}

// 创建默认服务
func NewServerBootstrap(opts ...Option) iface.Server {
	return newServerWithConfig(conf.DefaultConfig(), opts...)
}

// 自定义配置创建服务
func NewServerBootstrapWithConfig(config *conf.Config, opts ...Option) iface.Server {
	return newServerWithConfig(config, opts...)
}

func newServerWithConfig(config *conf.Config, opts ...Option) iface.Server {
	s := &ServerBootstrap{
		ip:      config.Host,
		port:    config.Port,
		decoder: nil,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *ServerBootstrap) Start() error {
	// 校验相关参数是否正常
	if err := s.verifyParam(); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	defer listener.Close()

	var cid uint32

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// 创建连接
		connection := NewConnection(conn, cid)
		cid++

		// 启动连接
		go connection.Start()
	}
}

func (s *ServerBootstrap) verifyParam() error {
	if s.addr == "" {
		return errors.New("addr must be required")
	}
	if s.decoder == nil {
		return errors.New("decoder must be required")
	}

	return nil
}

func (s *ServerBootstrap) Stop() {

}
