package bootstrap

import (
	"fmt"
	"math/rand"

	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/connection"
	"github.com/ningzining/lazynet/iface"
)

type Client struct {
	Bootstrap // 组合自基础启动器

	config *conf.ClientConfig // 配置
	conn   iface.Connection   // 连接对象
}

func NewClient(opts ...conf.ClientOption) iface.Client {
	return NewClientWithConfig(conf.DefaultClientConfig(), opts...)
}

// NewClientWithConfig 自定义配置创建服务
func NewClientWithConfig(config *conf.ClientConfig, opts ...conf.ClientOption) iface.Client {
	return newClientWithConfig(config, opts...)
}

// 使用配置创建服务
func newClientWithConfig(config *conf.ClientConfig, opts ...conf.ClientOption) iface.Client {
	for _, opt := range opts {
		opt(config)
	}

	c := &Client{
		Bootstrap: NewBootstrap(config.WorkerPoolSize, config.TaskQueueSize),
		config:    config,
		conn:      nil,
	}

	return c
}

// Start 启动客户端
func (c *Client) Start() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.config.Ip, c.config.Port))
	if err != nil {
		return err
	}

	log.Infof("tcp client link successfully localAddr: %s, remoteAddr: %s", conn.LocalAddr().String(), conn.RemoteAddr().String())

	// 开启工作线程池
	c.dispatcher.StartWorkerPool()

	newConnection := connection.New(c, conn, rand.Uint32(), c.config.MaxPackageSize)
	c.conn = newConnection

	newConnection.Start()

	return nil
}

func (c *Client) Stop() {
	c.conn.Stop()
}

func (c *Client) GetConfig() *conf.ClientConfig {
	return c.config
}

func (c *Client) GetConn() iface.Connection {
	return c.conn
}
