package iface

import "github.com/ningzining/lazynet/decoder"

type Server interface {
	Start() error
	Stop()
	SetDecoder(decoder decoder.Decoder)
	GetDecoder() decoder.Decoder
}
