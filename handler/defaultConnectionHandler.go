package handler

import (
	"fmt"
	"log"

	"github.com/ningzining/lazynet/iface"
)

type DefaultConnectionHandler struct {
}

func NewDefaultConnectionHandler() *DefaultConnectionHandler {
	return &DefaultConnectionHandler{}
}

func (d *DefaultConnectionHandler) ConnectionRead(ctx iface.Context, msg []byte) {
	log.Println(string(msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("server: %s", msg))); err != nil {
		return
	}
}
