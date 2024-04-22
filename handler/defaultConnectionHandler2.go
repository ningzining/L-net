package handler

import (
	"fmt"
	"log"

	"github.com/ningzining/lazynet/iface"
)

type DefaultConnectionHandler2 struct {
	BaseConnectionHandler
}

func NewDefaultConnectionHandler2() iface.ConnectionHandler {
	return &DefaultConnectionHandler2{}
}

func (d *DefaultConnectionHandler2) PreHandle(ctx iface.Context, msg []byte) {

}

func (d *DefaultConnectionHandler2) ChannelRead(ctx iface.Context, msg []byte) {
	log.Println(string(msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("server2: %s", msg))); err != nil {
		return
	}
	ctx.FireRead(msg)
}

func (d *DefaultConnectionHandler2) PostHandle(ctx iface.Context, msg []byte) {

}
