package handler

import (
	"fmt"
	"log"

	"github.com/ningzining/lazynet/iface"
)

type DefaultConnectionHandler struct {
	BaseConnectionHandler
}

func NewDefaultConnectionHandler() iface.ConnectionHandler {
	return &DefaultConnectionHandler{}
}

func (d *DefaultConnectionHandler) PreHandle(ctx iface.Context, msg []byte) {

}

func (d *DefaultConnectionHandler) ChannelRead(ctx iface.Context, msg []byte) {
	log.Println(fmt.Sprintf("client-%d: %s", ctx.GetConnection().GetConnID(), msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("%d-server1: %s", ctx.GetConnection().GetConnID(), msg))); err != nil {
		return
	}
	ctx.FireRead(msg)
}

func (d *DefaultConnectionHandler) PostHandle(ctx iface.Context, msg []byte) {

}
