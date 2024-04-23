package bootstrap

import (
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/server"
)

type Bootstrap struct {
	decoder decoder.Decoder // 解码器
	encoder encoder.Encoder // 编码器

	connOnActiveFunc func(conn iface.Connection) // 连接激活回调函数
	connOnCloseFunc  func(conn iface.Connection) // 连接关闭回调函数

	handlers []iface.ChannelHandler // 处理器

	dispatcher iface.Dispatcher // 消息分发器, 业务使用goroutine去处理
}

func NewBootstrap(workerPoolSize int, taskQueueSize int) Bootstrap {
	return Bootstrap{
		decoder:          nil,
		encoder:          nil,
		connOnActiveFunc: nil,
		connOnCloseFunc:  nil,
		handlers:         nil,
		dispatcher:       server.NewDispatcher(workerPoolSize, taskQueueSize),
	}
}

func (b *Bootstrap) Start() error {
	return nil
}

func (b *Bootstrap) Stop() {}

func (b *Bootstrap) SetEncoder(encoder encoder.Encoder) {
	b.encoder = encoder
}

func (b *Bootstrap) GetEncoder() encoder.Encoder {
	return b.encoder
}

func (b *Bootstrap) SetDecoder(decoder decoder.Decoder) {
	b.decoder = decoder
}

func (b *Bootstrap) GetDecoder() decoder.Decoder {
	return b.decoder
}

func (b *Bootstrap) SetConnOnActiveFunc(f func(conn iface.Connection)) {
	b.connOnActiveFunc = f
}

func (b *Bootstrap) GetConnOnActiveFunc() func(conn iface.Connection) {
	return b.connOnActiveFunc
}

func (b *Bootstrap) SetConnOnCloseFunc(f func(conn iface.Connection)) {
	b.connOnCloseFunc = f
}

func (b *Bootstrap) GetConnOnCloseFunc() func(conn iface.Connection) {
	return b.connOnCloseFunc
}

func (b *Bootstrap) AddChannelHandler(handler iface.ChannelHandler) {
	b.handlers = append(b.handlers, handler)
}

func (b *Bootstrap) GetChannelHandlers() []iface.ChannelHandler {
	return b.handlers
}

func (b *Bootstrap) GetDispatcher() iface.Dispatcher {
	return b.dispatcher
}

var _ iface.Bootstrap = &Bootstrap{}
