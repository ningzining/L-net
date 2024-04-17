package pipeline

import (
	"context"

	"github.com/ningzining/lazynet/iface"
)

type ConnectionContext struct {
	context.Context

	pipeline iface.Pipeline
	handler  iface.ConnectionHandler

	prev *ConnectionContext
	next *ConnectionContext
}

func NewContext(ctx context.Context, pipeline iface.Pipeline, handler iface.ConnectionHandler) *ConnectionContext {
	return &ConnectionContext{
		Context:  ctx,
		pipeline: pipeline,
		handler:  handler,
	}
}

func (c *ConnectionContext) GetHandler() iface.ConnectionHandler {
	return c.handler
}

func (c *ConnectionContext) GetConnection() iface.Connection {
	return c.pipeline.GetConnection()
}

func (c *ConnectionContext) FireConnectionRead(msg []byte) {
	if c.next == nil || c.next.GetHandler() == nil {
		return
	}
	c.next.GetHandler().PreHandle(c.next, msg)
	c.next.GetHandler().ConnectionRead(c.next, msg)
	c.next.GetHandler().PostHandle(c.next, msg)
}
