package bootstrap

import (
	"bytes"
	"context"
	"net"

	"github.com/ningzining/lazynet/decoder"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/handler"
	"github.com/ningzining/lazynet/iface"
)

type Connection struct {
	conn       net.Conn
	connID     uint32
	remoteAddr net.Addr
	localAddr  net.Addr

	server     iface.Server
	decoder    decoder.Decoder
	msgHandler handler.ConnectionHandler

	readBuffer  *bytes.Buffer // 读取缓冲区
	writeBuffer *bytes.Buffer // 写入缓冲区

	onActive func(conn iface.Connection) // 钩子函数，当连接建立的时候调用
	onClose  func(conn iface.Connection) // 钩子函数，当连接断开的时候调用
}

func NewConnection(server iface.Server, conn net.Conn, connID uint32) iface.Connection {
	return &Connection{
		server:      server,
		conn:        conn,
		connID:      connID,
		decoder:     server.GetDecoder(),
		remoteAddr:  conn.RemoteAddr(),
		localAddr:   conn.LocalAddr(),
		onActive:    server.GetConnOnActiveFunc(),
		onClose:     server.GetConnOnCloseFunc(),
		msgHandler:  server.GetMsgHandler(),
		readBuffer:  bytes.NewBuffer(make([]byte, 0, 4096)),
		writeBuffer: bytes.NewBuffer(make([]byte, 0, 4096)),
	}
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
		readBytes := make([]byte, 1024)
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
				c.msgHandler.ConnectionRead(context.Background(), frame)
			}
		} else {
			c.msgHandler.ConnectionRead(context.Background(), c.readBuffer.Bytes())
			c.readBuffer.Reset()
		}
	}

	// 执行连接关闭的钩子函数
	c.callOnClose()
}

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

var _ iface.Connection = &Connection{}
