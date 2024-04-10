package iface

import (
	"github.com/ningzining/lazynet/encoder"
)

type Client interface {
	Start() error
	Stop()

	Read() ([]byte, error)
	Write(msg []byte) error

	SetEncoder(decoder encoder.Encoder)
	GetEncoder() encoder.Encoder
}
