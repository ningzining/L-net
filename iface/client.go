package iface

import (
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
)

type Client interface {
	Start() error // 启动客户端
	Stop()        // 停止客户端

	Write(msg []byte) error // 写入数据

	SetEncoder(decoder encoder.Encoder) // 设置编码器
	GetEncoder() encoder.Encoder        // 获取编码器
	SetDecoder(decoder decoder.Decoder) // 设置解码器
	GetDecoder() decoder.Decoder        // 获取解码器

	AddChannelHandler(handler ChannelHandler) // 添加处理器
	GetChannelHandlers() []ChannelHandler     // 获取处理器

	SetConnOnActiveFunc(func(conn Connection))  // 设置连接激活的回调函数
	GetConnOnActiveFunc() func(conn Connection) // 获取连接激活的回调函数
	SetConnOnCloseFunc(func(conn Connection))   // 设置连接关闭的回调函数
	GetConnOnCloseFunc() func(conn Connection)  // 获取连接关闭的回调函数

	GetDispatcher() Dispatcher // 获取消息分发器
}
