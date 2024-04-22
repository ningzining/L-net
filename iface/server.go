package iface

import (
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
)

type Server interface {
	Start() error // 启动服务
	Stop()        // 关闭服务

	GetConfig() *conf.Config // 获取当前服务的配置

	SetDecoder(decoder decoder.Decoder) // 设置解码器
	GetDecoder() decoder.Decoder        // 获取解码器
	SetEncoder(decoder encoder.Encoder) // 设置编码器
	GetEncoder() encoder.Encoder        // 获取编码器

	AddChannelHandler(handler ConnectionHandler) // 添加处理器
	GetChannelHandlers() []ConnectionHandler     // 获取处理器

	SetConnOnActiveFunc(func(conn Connection))  // 设置连接激活的回调函数
	GetConnOnActiveFunc() func(conn Connection) // 获取连接激活的回调函数
	SetConnOnCloseFunc(func(conn Connection))   // 设置连接关闭的回调函数
	GetConnOnCloseFunc() func(conn Connection)  // 获取连接关闭的回调函数

	GetConnManager() ConnManager // 获取连接管理器
	GetDispatcher() Dispatcher   // 获取消息分发器
}
