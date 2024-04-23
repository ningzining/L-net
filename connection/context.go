package connection

import (
	"context"

	"github.com/ningzining/lazynet/iface"
)

type ChannelContext struct {
	context.Context

	pipeline iface.Pipeline
	handler  iface.ChannelHandler

	prev *ChannelContext
	next *ChannelContext
}

func NewContext(ctx context.Context, pipeline iface.Pipeline, handler iface.ChannelHandler) *ChannelContext {
	return &ChannelContext{
		Context:  ctx,
		pipeline: pipeline,
		handler:  handler,
	}
}

func (c *ChannelContext) GetHandler() iface.ChannelHandler {
	return c.handler
}

func (c *ChannelContext) GetConnection() iface.Connection {
	return c.pipeline.GetConnection()
}

// FireRead 调用链表当中的下一个处理器
func (c *ChannelContext) FireRead(msg []byte) {
	if c.next == nil || c.next.GetHandler() == nil {
		return
	}

	c.next.DoHandle(msg)
}

// DoHandle 调用当前处理器处理消息
func (c *ChannelContext) DoHandle(msg []byte) {
	c.handler.PreHandle(c, msg)
	c.handler.ChannelRead(c, msg)
	c.handler.PostHandle(c, msg)
}
