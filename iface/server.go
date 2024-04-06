package iface

import (
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/handler"
)

type Server interface {
	Start() error
	Stop()

	SetDecoder(decoder decoder.Decoder)
	GetDecoder() decoder.Decoder

	SetMsgHandler(handler handler.ConnectionHandler)
	GetMsgHandler() handler.ConnectionHandler

	SetConnOnActiveFunc(func(conn Connection))
	GetConnOnActiveFunc() func(conn Connection)

	SetConnOnCloseFunc(func(conn Connection))
	GetConnOnCloseFunc() func(conn Connection)
}
