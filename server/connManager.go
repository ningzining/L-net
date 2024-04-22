package server

import (
	"sync"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/iface"
)

type ConnManager struct {
	connections map[uint32]iface.Connection // 连接集合
	connLock    sync.RWMutex                // 读写锁
}

func NewConnManager() iface.ConnManager {
	return &ConnManager{
		connections: make(map[uint32]iface.Connection),
	}
}

func (c *ConnManager) Add(conn iface.Connection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	log.Infof("connection[%d] add successfully, conn count: %d", conn.GetConnID(), c.Size())
}

func (c *ConnManager) Remove(connId uint32) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, connId)
	log.Infof("connection[%d] remove successfully, conn count: %d", connId, c.Size())
}

func (c *ConnManager) Get(connId uint32) (iface.Connection, bool) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	connection, ok := c.connections[connId]

	return connection, ok
}

func (c *ConnManager) Size() int {
	return len(c.connections)
}

func (c *ConnManager) Clear() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connId, connection := range c.connections {
		connection.Stop()
		delete(c.connections, connId)
	}

	log.Info("clear all connections successfully")
}
