package iface

import (
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
)

type Server interface {
	Start() error
	Stop()

	GetConfig() *conf.Config

	SetDecoder(decoder decoder.Decoder)
	GetDecoder() decoder.Decoder

	AddConnectionHandler(handler ConnectionHandler)
	GetConnectionHandler() ConnectionHandler

	SetConnOnActiveFunc(func(conn Connection))
	GetConnOnActiveFunc() func(conn Connection)

	SetConnOnCloseFunc(func(conn Connection))
	GetConnOnCloseFunc() func(conn Connection)
}
