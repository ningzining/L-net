package request

import (
	"github.com/ningzining/lazynet/iface"
)

type Request struct {
	conn iface.Connection
	msg  []byte
}

func NewRequest(conn iface.Connection, msg []byte) iface.Request {
	return &Request{
		conn: conn,
		msg:  msg,
	}
}

func (r *Request) GetConn() iface.Connection {
	return r.conn
}

func (r *Request) GetMsg() []byte {
	return r.msg
}
