package bootstrap

import (
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
)

type Bootstrap struct {
	decoder decoder.Decoder // 解码器
	encoder encoder.Encoder // 编码器

	handlers []iface.ChannelHandler // 处理器

	dispatcher iface.Dispatcher // 消息分发器,业务使用goroutine去处理

	connOnActiveFunc func(conn iface.Connection) // 连接激活回调函数
	connOnCloseFunc  func(conn iface.Connection) // 连接关闭回调函数
}
