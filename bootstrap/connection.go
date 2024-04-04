package bootstrap

import (
	"context"
	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/handler"
	"github.com/ningzining/lazynet/iface"
)

type Connection struct {
	connID     uint32
	conn       net.Conn
	remoteAddr net.Addr
	localAddr  net.Addr
	handler    handler.ConnectionHandler
}

func NewConnection(conn net.Conn, connID uint32) iface.Connection {
	return &Connection{
		connID:     connID,
		conn:       conn,
		remoteAddr: conn.RemoteAddr(),
		localAddr:  conn.LocalAddr(),
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

func (c *Connection) Read(b []byte) (n int, err error) {
	return c.Conn().Read(b)
}

func (c *Connection) Write(b []byte) (n int, err error) {
	return c.Conn().Write(b)
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
		_, err := c.conn.Read(readBytes)
		if err != nil {
			continue
		}
		// 处理数据
		c.handler.ConnectionRead(context.Background(), readBytes)
	}

}
func (c *Connection) Stop() {
	c.conn.Close()
}

var _ iface.Connection = &Connection{}
