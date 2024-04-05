package bootstrap

import (
	"context"
	"github.com/ningzining/lazynet/decoder"
	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/handler"
	"github.com/ningzining/lazynet/iface"
)

type Connection struct {
	server     iface.Server
	conn       net.Conn
	connID     uint32
	decoder    decoder.Decoder
	remoteAddr net.Addr
	localAddr  net.Addr
	handler    handler.ConnectionHandler
}

func NewConnection(server iface.Server, conn net.Conn, connID uint32) iface.Connection {
	return &Connection{
		server:     server,
		decoder:    server.GetDecoder(),
		conn:       conn,
		connID:     connID,
		remoteAddr: conn.RemoteAddr(),
		localAddr:  conn.LocalAddr(),
		handler:    handler.NewDefaultConnectionHandler(),
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
			continue
		}
		// 处理数据
		c.handler.ConnectionRead(context.Background(), readBytes[:n])
	}

}

func (c *Connection) Stop() {
	c.conn.Close()
}

var _ iface.Connection = &Connection{}
