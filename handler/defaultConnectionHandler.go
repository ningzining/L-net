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

func (d *DefaultConnectionHandler) ConnectionRead(ctx iface.Context, msg []byte) {
	log.Println(fmt.Sprintf("client-%d: %s", ctx.GetConnection().ConnID(), msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("%d-server1: %s", ctx.GetConnection().ConnID(), msg))); err != nil {
		return
	}
	ctx.FireConnectionRead(msg)
}

func (d *DefaultConnectionHandler) PostHandle(ctx iface.Context, msg []byte) {

}
