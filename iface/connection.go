package iface

import (
	"net"
)

type Connection interface {
	Conn() net.Conn
	ConnID() uint32

	RemoteAddr() net.Addr
	LocalAddr() net.Addr

	Start()
	Stop()

	Write(msg []byte) error
}
