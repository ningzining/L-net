package handler

import (
	"github.com/ningzining/lazynet/iface"
)

// BaseChannelHandler
// 定义基础的连接处理器，实现ConnectionHandler接口
// 如果要实现自定义的连接处理器，需要继承这个类，并选择实现PreHandle、ChannelRead、PostHandle方法
type BaseChannelHandler struct {
}

func NewBaseConnectionHandler() iface.ChannelHandler {
	return &BaseChannelHandler{}
}

func (d *BaseChannelHandler) PreHandle(ctx iface.Context, msg []byte) {}

func (d *BaseChannelHandler) ChannelRead(ctx iface.Context, msg []byte) {}

func (d *BaseChannelHandler) PostHandle(ctx iface.Context, msg []byte) {}
