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
	log.Println(string(msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("server: %s", msg))); err != nil {
		return
	}
}

func (d *DefaultConnectionHandler) PostHandle(ctx iface.Context, msg []byte) {

}
