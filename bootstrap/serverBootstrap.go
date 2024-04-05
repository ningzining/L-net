package bootstrap

import (
	"fmt"
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

// 使用配置创建服务
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
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
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
		connection := NewConnection(s, conn, cid)
		cid++

		// 启动连接
		go connection.Start()
	}
}

func (s *ServerBootstrap) Stop() {

}

func (s *ServerBootstrap) SetDecoder(decoder decoder.Decoder) {
	s.decoder = decoder
}

func (s *ServerBootstrap) GetDecoder() decoder.Decoder {
	if s.decoder == nil {
		s.decoder = decoder.NewLineBasedFrameDecoder()
	}

	return s.decoder
}
