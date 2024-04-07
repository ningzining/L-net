package handler

import (
	"context"

	"github.com/ningzining/lazynet/iface"
)

type ConnectionContext struct {
	context.Context
	connection iface.Connection
}

func NewConnectionContext(ctx context.Context, conn iface.Connection) iface.Context {
	return &ConnectionContext{
		Context:    ctx,
		connection: conn,
	}
}

func (c *ConnectionContext) GetConnection() iface.Connection {
	return c.connection
}
