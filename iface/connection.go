package iface

import (
	"net"
)

type Connection interface {
	Conn() net.Conn
	ConnID() uint32
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Start()
	Stop()
}
