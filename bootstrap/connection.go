package bootstrap

import (
	"bytes"
	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/pipeline"
)

type Connection struct {
	conn       net.Conn
	connID     uint32
	remoteAddr net.Addr
	localAddr  net.Addr

	server  iface.Server
	decoder decoder.Decoder // 解码器
	encoder encoder.Encoder // 编码器

	pipeline iface.Pipeline // 处理器管道

	readBuffer  *bytes.Buffer // 读取缓冲区
	writeBuffer *bytes.Buffer // 写入缓冲区

	onActive func(conn iface.Connection) // 钩子函数，当连接建立的时候调用
	onClose  func(conn iface.Connection) // 钩子函数，当连接断开的时候调用
}

func NewConnection(server iface.Server, conn net.Conn, connID uint32) iface.Connection {
	c := &Connection{
		server:      server,
		conn:        conn,
		connID:      connID,
		decoder:     server.GetDecoder(),
		encoder:     server.GetEncoder(),
		remoteAddr:  conn.RemoteAddr(),
		localAddr:   conn.LocalAddr(),
		onActive:    server.GetConnOnActiveFunc(),
		onClose:     server.GetConnOnCloseFunc(),
		pipeline:    nil,
		readBuffer:  bytes.NewBuffer(make([]byte, 0, server.GetConfig().MaxPackageSize*4)),
		writeBuffer: bytes.NewBuffer(make([]byte, 0, server.GetConfig().MaxPackageSize*4)),
	}

	c.pipeline = pipeline.NewPipeline(c)
	for _, handler := range server.GetConnectionHandlers() {
		c.pipeline.AddLast(handler)
	}

	return c
}

func (c *Connection) ConnID() uint32 {
	return c.connID
}

func (c *Connection) Conn() net.Conn {
	return c.conn
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.remoteAddr
}

func (c *Connection) LocalAddr() net.Addr {
	return c.localAddr
}

func (c *Connection) Start() {
	// 执行连接建立的钩子函数
	c.callOnActive()

	// 全双工通信，可以接收数据也可以写入数据
	// 启动阅读器
	go c.StartReader()
	// 启动写入器
	go c.StartWriter()
}

func (c *Connection) StartReader() {
	defer func() {
		c.Stop()
		if err := recover(); err != nil {
			log.Errorf("%v", err)
		}
	}()

	for {
		// 读取数据
		readBytes := make([]byte, c.server.GetConfig().MaxPackageSize)
		n, err := c.conn.Read(readBytes)
		if err != nil {
			break
		}
		// 写入连接的缓冲区
		c.readBuffer.Write(readBytes[:n])
		// 使用注册的解码器进行解码
		if c.decoder != nil {
			// 一个数据包可能包含多个数据帧的情况，所以需要循环处理
			frames := c.decoder.Decode(c.readBuffer)
			// 读取每一帧的数据并进行处理
			for _, frame := range frames {
				c.pipeline.Handle(frame)
			}
		} else {
			c.pipeline.Handle(c.readBuffer.Bytes())
			c.readBuffer.Reset()
		}
	}

	// 执行连接关闭的钩子函数
	c.callOnClose()
}

// StartWriter
// todo: 实现写入器
func (c *Connection) StartWriter() {

}

func (c *Connection) callOnActive() {
	if c.onActive != nil {
		c.onActive(c)
	}
}

func (c *Connection) callOnClose() {
	if c.onClose != nil {
		c.onClose(c)
	}
}

func (c *Connection) Stop() {
	c.conn.Close()
}

func (c *Connection) Write(msg []byte) error {
	res := msg
	if c.encoder != nil {
		var err error
		if res, err = c.encoder.Encode(res); err != nil {
			return err
		}
	}

	if _, err := c.conn.Write(res); err != nil {
		return err
	}

	return nil
}

var _ iface.Connection = &Connection{}
