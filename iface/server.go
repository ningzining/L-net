package iface

import (
	"github.com/ningzining/lazynet/decoder"
)

type Server interface {
	Start() error
	Stop()

	SetDecoder(decoder decoder.Decoder)
	GetDecoder() decoder.Decoder

	SetConnectionHandler(handler ConnectionHandler)
	GetConnectionHandler() ConnectionHandler

	SetConnOnActiveFunc(func(conn Connection))
	GetConnOnActiveFunc() func(conn Connection)

	SetConnOnCloseFunc(func(conn Connection))
	GetConnOnCloseFunc() func(conn Connection)
}
