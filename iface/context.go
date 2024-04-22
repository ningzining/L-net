package iface

import (
	"context"
)

type Context interface {
	context.Context

	GetHandler() ChannelHandler // 获取当前的处理器
	GetConnection() Connection  // 获取连接对象

	DoHandle(msg []byte) // 处理义务逻辑
	FireRead(msg []byte) // 调用下一个handler处理业务
}
