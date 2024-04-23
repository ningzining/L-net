package bootstrap

import (
	"errors"
	"fmt"
	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/connection"
	"github.com/ningzining/lazynet/iface"
)

type Server struct {
	Bootstrap // 组合自基础启动器

	config      *conf.ServerConfig // 配置
	connManager iface.ConnManager  // 连接管理器
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
		Bootstrap:   NewBootstrap(config.WorkerPoolSize, config.TaskQueueSize),
		config:      config,
		connManager: connection.NewConnManager(),
	}

	return s
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
			conn.Close()
			continue
		}

		// 创建连接
		newConnection := connection.New(s, conn, cid, s.config.MaxPackageSize)
		s.connManager.Add(newConnection)

		cid++

		// 启动连接
		go newConnection.Start()
	}
}

func (s *Server) verify() error {
	if len(s.Bootstrap.GetChannelHandlers()) == 0 {
		return errors.New("connectionHandler must be added")
	}

	return nil
}

func (s *Server) Stop() {
	// 释放资源
	// 回收所有的连接
	s.connManager.Clear()

	log.Infof("tcp server stop successfully at: %s:%d", s.config.Host, s.config.Port)
}

func (s *Server) GetConfig() *conf.ServerConfig {
	return s.config
}

func (s *Server) GetConnManager() iface.ConnManager {
	return s.connManager
}
