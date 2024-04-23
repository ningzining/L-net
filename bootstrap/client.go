package bootstrap

import (
	"bytes"
	"fmt"

	"net"

	"github.com/ningzining/lazynet/client"
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/server"
)

type Client struct {
	config *conf.ClientConfig

	conn iface.Connection // 连接对象

	encoder encoder.Encoder // 编码器
	decoder decoder.Decoder // 解码器

	readBuffer *bytes.Buffer // 读取缓冲区

	handlerList []iface.ChannelHandler
	pipeline    iface.Pipeline

	dispatcher iface.Dispatcher // 消息分发器,业务使用goroutine去处理

	connOnActiveFunc func(conn iface.Connection)
	connOnCloseFunc  func(conn iface.Connection)

	exitChan chan struct{}
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
		config:     config,
		conn:       nil,
		encoder:    nil,
		decoder:    nil,
		readBuffer: bytes.NewBuffer(make([]byte, 0, 1024)),
		pipeline:   nil,
		dispatcher: server.NewDispatcher(config.WorkerPoolSize, config.TaskQueueSize), handlerList: nil,
		connOnActiveFunc: nil,
		connOnCloseFunc:  nil,
	}

	return c
}

func (c *Client) GetConfig() *conf.ClientConfig {
	return c.config
}

// Start 启动客户端
func (c *Client) Start() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.config.Ip, c.config.Port))
	if err != nil {
		return err
	}

	// 开启工作线程池
	c.dispatcher.StartWorkerPool()

	connection := client.NewConnection(c, conn)
	c.conn = connection

	connection.Start()

	return nil
}

func (c *Client) Stop() {
	c.exitChan <- struct{}{}

	close(c.exitChan)
}

// SetEncoder 设置编码器
func (c *Client) SetEncoder(encoder encoder.Encoder) {
	c.encoder = encoder
}

// GetEncoder 获取编码器
func (c *Client) GetEncoder() encoder.Encoder {
	return c.encoder
}

// SetDecoder 设置解码器
func (c *Client) SetDecoder(decoder decoder.Decoder) {
	c.decoder = decoder
}

// GetDecoder 获取解码器
func (c *Client) GetDecoder() decoder.Decoder {
	return c.decoder
}

func (c *Client) AddChannelHandler(handler iface.ChannelHandler) {
	c.handlerList = append(c.handlerList, handler)
}

func (c *Client) GetChannelHandlers() []iface.ChannelHandler {
	return c.handlerList
}

func (c *Client) SetConnOnActiveFunc(f func(conn iface.Connection)) {
	c.connOnActiveFunc = f
}

func (c *Client) GetConnOnActiveFunc() func(conn iface.Connection) {
	return c.connOnActiveFunc
}

func (c *Client) SetConnOnCloseFunc(f func(conn iface.Connection)) {
	c.connOnCloseFunc = f
}

func (c *Client) GetConnOnCloseFunc() func(conn iface.Connection) {
	return c.connOnCloseFunc
}

func (c *Client) GetDispatcher() iface.Dispatcher {
	return c.dispatcher
}

func (c *Client) GetConn() iface.Connection {
	return c.conn
}
