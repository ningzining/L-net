package bootstrap

import (
	"bytes"

	"net"

	"github.com/ningzining/lazynet/client"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/server"
)

type Client struct {
	addr string           // 地址
	conn iface.Connection // 连接对象

	encoder encoder.Encoder // 编码器
	decoder decoder.Decoder // 解码器

	readBuffer *bytes.Buffer // 读取缓冲区

	handlerList []iface.ChannelHandler
	pipeline    iface.Pipeline

	dispatcher iface.Dispatcher // 消息分发器,业务使用goroutine去处理

	connOnActiveFunc func(conn iface.Connection)
	connOnCloseFunc  func(conn iface.Connection)
}

func NewClient(addr string) *Client {
	return &Client{
		addr:       addr,
		readBuffer: bytes.NewBuffer(make([]byte, 0, 1024)),
		dispatcher: server.NewDispatcher(4, 1024),
	}
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

// Start 启动客户端
func (c *Client) Start() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.dispatcher.StartWorkerPool()
	connection := client.NewConnection(c, conn)
	c.conn = connection
	connection.Start()
	return nil
}

func (c *Client) Stop() {
	c.conn.Stop()
}

// 往连接中写入字节数组
func (c *Client) Write(source []byte) error {
	return c.conn.Write(source)
}

func (c *Client) AddChannelHandler(handler iface.ChannelHandler) {
	c.handlerList = append(c.handlerList, handler)
}

func (c *Client) GetChannelHandlers() []iface.ChannelHandler {
	return c.handlerList
}

func (c *Client) GetDispatcher() iface.Dispatcher {
	return c.dispatcher
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
