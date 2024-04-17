package handler

import (
	"fmt"
	"log"

	"github.com/ningzining/lazynet/iface"
)

type DefaultConnectionHandler3 struct {
	BaseConnectionHandler
}

func NewDefaultConnectionHandler3() iface.ConnectionHandler {
	return &DefaultConnectionHandler3{}
}

func (d *DefaultConnectionHandler3) PreHandle(ctx iface.Context, msg []byte) {

}

func (d *DefaultConnectionHandler3) ConnectionRead(ctx iface.Context, msg []byte) {
	log.Println(string(msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("server3: %s", msg))); err != nil {
		return
	}
}

func (d *DefaultConnectionHandler3) PostHandle(ctx iface.Context, msg []byte) {

}
