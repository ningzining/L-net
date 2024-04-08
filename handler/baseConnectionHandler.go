package handler

import (
	"github.com/ningzining/lazynet/iface"
)

// BaseConnectionHandler
// 定义基础的连接处理器，实现ConnectionHandler接口
// 如果要实现自定义的连接处理器，需要继承这个类，并选择实现PreHandle、ConnectionRead、PostHandle方法
type BaseConnectionHandler struct {
}

var _ iface.ConnectionHandler = &BaseConnectionHandler{}

func (d *BaseConnectionHandler) PreHandle(ctx iface.Context, msg []byte) {}

func (d *BaseConnectionHandler) ConnectionRead(ctx iface.Context, msg []byte) {}

func (d *BaseConnectionHandler) PostHandle(ctx iface.Context, msg []byte) {}
