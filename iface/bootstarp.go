package iface

import (
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
)

// Bootstrap 启动器
type Bootstrap interface {
	Start() error // 启动服务
	Stop()        // 停止服务

	SetEncoder(encoder encoder.Encoder) // 设置编码器
	GetEncoder() encoder.Encoder        // 获取编码器
	SetDecoder(decoder decoder.Decoder) // 设置解码器
	GetDecoder() decoder.Decoder        // 获取解码器

	AddChannelHandler(handler ChannelHandler) // 添加处理器
	GetChannelHandlers() []ChannelHandler     // 获取处理器

	GetDispatcher() Dispatcher // 获取消息分发器

	SetConnOnActiveFunc(func(conn Connection))  // 设置连接激活的回调函数
	GetConnOnActiveFunc() func(conn Connection) // 获取连接激活的回调函数
	SetConnOnCloseFunc(func(conn Connection))   // 设置连接关闭的回调函数
	GetConnOnCloseFunc() func(conn Connection)  // 获取连接关闭的回调函数
}

type Client interface {
	Bootstrap

	GetConfig() *conf.ClientConfig // 获取当前服务的配置

	GetConn() Connection // 获取连接

}

type Server interface {
	Bootstrap

	GetConfig() *conf.ServerConfig // 获取当前服务的配置

	GetConnManager() ConnManager // 获取连接管理器
}
