package bootstrap

import (
	"bytes"
	"context"
	"github.com/ningzining/lazynet/decoder"
	"net"

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

	buffer *bytes.Buffer

	onActive func(conn iface.Connection) // 钩子函数，当连接建立的时候调用
	onClose  func(conn iface.Connection) // 钩子函数，当连接断开的时候调用
}

func NewConnection(server iface.Server, conn net.Conn, connID uint32) iface.Connection {
	return &Connection{
		server:     server,
		conn:       conn,
		connID:     connID,
		decoder:    server.GetDecoder(),
		remoteAddr: conn.RemoteAddr(),
		localAddr:  conn.LocalAddr(),
		onActive:   server.GetConnOnActiveFunc(),
		onClose:    server.GetConnOnCloseFunc(),
		msgHandler: server.GetMsgHandler(),
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
	if c.onActive != nil {
		c.onActive(c)
	}

	// 启动阅读器
	go c.StartReader()
	// todo: 启动写入器
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
		// todo: 写入连接的缓冲区后进行处理数据
		// 处理数据
		c.msgHandler.ConnectionRead(context.Background(), readBytes[:n])
	}

	if c.onClose != nil {
		c.onClose(c)
	}
}

func (c *Connection) Stop() {
	c.conn.Close()
}

var _ iface.Connection = &Connection{}
