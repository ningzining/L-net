package iface

import (
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
)

type Server interface {
	Start() error
	Stop()

	GetConfig() *conf.Config

	SetDecoder(decoder decoder.Decoder)
	GetDecoder() decoder.Decoder

	SetEncoder(decoder encoder.Encoder)
	GetEncoder() encoder.Encoder

	AddConnectionHandler(handler ConnectionHandler)
	GetConnectionHandlers() []ConnectionHandler

	SetConnOnActiveFunc(func(conn Connection))
	GetConnOnActiveFunc() func(conn Connection)

	SetConnOnCloseFunc(func(conn Connection))
	GetConnOnCloseFunc() func(conn Connection)

	GetConnManager() ConnManager
	GetDispatcher() Dispatcher
}
