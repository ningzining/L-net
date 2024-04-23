package bootstrap

import (
	"errors"
	"fmt"
	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/server"
)

type Server struct {
	config *conf.ServerConfig

	decoder decoder.Decoder // 解码器
	encoder encoder.Encoder // 编码器

	handlers []iface.ChannelHandler

	connManager iface.ConnManager
	// 消息分发器,业务使用goroutine去处理
	dispatcher iface.Dispatcher

	connOnActiveFunc func(conn iface.Connection)
	connOnCloseFunc  func(conn iface.Connection)
}

// NewServer 创建默认服务
func NewServer(opts ...conf.ServerOption) iface.Server {
	return NewServerWithConfig(conf.DefaultServerConfig(), opts...)
}

// NewServerWithConfig 自定义配置创建服务
func NewServerWithConfig(config *conf.ServerConfig, opts ...conf.ServerOption) iface.Server {
	return newServerWithConfig(config, opts...)
}

// 使用配置创建服务
func newServerWithConfig(config *conf.ServerConfig, opts ...conf.ServerOption) iface.Server {
	for _, opt := range opts {
		opt(config)
	}

	s := &Server{
		config:           config,
		decoder:          nil,
		encoder:          nil,
		handlers:         make([]iface.ChannelHandler, 0),
		connOnActiveFunc: nil,
		connOnCloseFunc:  nil,
		connManager:      server.NewConnManager(),
		dispatcher:       server.NewDispatcher(config.WorkerPoolSize, config.TaskQueueSize),
	}

	return s
}

func (s *Server) GetConfig() *conf.ServerConfig {
	return s.config
}

func (s *Server) Start() error {
	if err := s.verify(); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Infof("tcp server listen at: %s", listener.Addr().String())
	s.dispatcher.StartWorkerPool()
	var cid uint32
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		if s.connManager.Size() >= s.config.MaxConnSize {
			log.Errorf("tcp server connection pool is full, max size: %d", s.config.MaxConnSize)
			// todo: 返回超过最大连接数错误
			conn.Close()
			continue
		}

		// 创建连接
		newConnection := server.NewConnection(s, conn, cid)
		cid++

		// 启动连接
		go newConnection.Start()
	}
}

func (s *Server) Stop() {
	// 释放资源
	// 回收所有的连接
	s.connManager.Clear()

	log.Infof("tcp server stop successfully at: %s:%d", s.config.Host, s.config.Port)
}

func (s *Server) verify() error {
	if len(s.handlers) == 0 {
		return errors.New("connectionHandler must be added")
	}

	return nil
}

func (s *Server) SetDecoder(decoder decoder.Decoder) {
	s.decoder = decoder
}

func (s *Server) GetDecoder() decoder.Decoder {
	return s.decoder
}

func (s *Server) SetEncoder(encoder encoder.Encoder) {
	s.encoder = encoder
}

func (s *Server) GetEncoder() encoder.Encoder {
	return s.encoder
}

func (s *Server) AddChannelHandler(handler iface.ChannelHandler) {
	s.handlers = append(s.handlers, handler)
}

func (s *Server) GetChannelHandlers() []iface.ChannelHandler {
	return s.handlers
}

func (s *Server) SetConnOnActiveFunc(f func(conn iface.Connection)) {
	s.connOnActiveFunc = f
}

func (s *Server) GetConnOnActiveFunc() func(conn iface.Connection) {
	return s.connOnActiveFunc
}

func (s *Server) SetConnOnCloseFunc(f func(conn iface.Connection)) {
	s.connOnCloseFunc = f
}

func (s *Server) GetConnOnCloseFunc() func(conn iface.Connection) {
	return s.connOnCloseFunc
}

func (s *Server) GetConnManager() iface.ConnManager {
	return s.connManager
}

func (s *Server) GetDispatcher() iface.Dispatcher {
	return s.dispatcher
}
