package bootstrap

import (
	"fmt"
	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/pipeline"
)

type ServerBootstrap struct {
	config conf.Config

	decoder decoder.Decoder // 解码器
	encoder encoder.Encoder // 编码器

	pipeline iface.Pipeline // 处理器管道

	connOnActiveFunc func(conn iface.Connection)
	connOnCloseFunc  func(conn iface.Connection)
}

// NewServerBootstrap 创建默认服务
func NewServerBootstrap(opts ...Option) iface.Server {
	return newServerWithConfig(conf.DefaultConfig(), opts...)
}

// NewServerBootstrapWithConfig 自定义配置创建服务
func NewServerBootstrapWithConfig(config *conf.Config, opts ...Option) iface.Server {
	return newServerWithConfig(config, opts...)
}

// 使用配置创建服务
func newServerWithConfig(config *conf.Config, opts ...Option) iface.Server {
	s := &ServerBootstrap{
		config: conf.Config{
			Host:           config.Host,
			Port:           config.Port,
			MaxPackageSize: config.MaxPackageSize,
		},
		decoder:          nil,
		pipeline:         pipeline.NewPipeline(nil),
		connOnActiveFunc: nil,
		connOnCloseFunc:  nil,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *ServerBootstrap) Start() error {
	if err := s.verify(); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Infof("tcp server listen at: %s", listener.Addr().String())

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

func (s *ServerBootstrap) verify() error {
	//if s.connectionHandler == nil {
	//	return errors.NewContext("connectionHandler must be added")
	//}

	return nil
}

func (s *ServerBootstrap) GetConfig() *conf.Config {
	return &s.config
}

func (s *ServerBootstrap) SetDecoder(decoder decoder.Decoder) {
	s.decoder = decoder
}

func (s *ServerBootstrap) GetDecoder() decoder.Decoder {
	return s.decoder
}

func (s *ServerBootstrap) SetEncoder(encoder encoder.Encoder) {
	s.encoder = encoder
}

func (s *ServerBootstrap) GetEncoder() encoder.Encoder {
	return s.encoder
}

func (s *ServerBootstrap) AddConnectionHandler(handler iface.ConnectionHandler) {
	s.pipeline.AddLast(handler)
}

func (s *ServerBootstrap) GetPipeline() iface.Pipeline {
	return s.pipeline
}

func (s *ServerBootstrap) SetConnOnActiveFunc(f func(conn iface.Connection)) {
	s.connOnActiveFunc = f
}

func (s *ServerBootstrap) GetConnOnActiveFunc() func(conn iface.Connection) {
	return s.connOnActiveFunc
}

func (s *ServerBootstrap) SetConnOnCloseFunc(f func(conn iface.Connection)) {
	s.connOnCloseFunc = f
}

func (s *ServerBootstrap) GetConnOnCloseFunc() func(conn iface.Connection) {
	return s.connOnCloseFunc
}
