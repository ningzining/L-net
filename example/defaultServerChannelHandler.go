package example

import (
	"fmt"
	"log"

	"github.com/ningzining/lazynet/handler"
	"github.com/ningzining/lazynet/iface"
)

type DefaultServerChannelHandler struct {
	handler.BaseChannelHandler
}

func NewDefaultServerChannelHandler() iface.ChannelHandler {
	return &DefaultServerChannelHandler{}
}

func (d *DefaultServerChannelHandler) PreHandle(ctx iface.Context, msg []byte) {

}

func (d *DefaultServerChannelHandler) ChannelRead(ctx iface.Context, msg []byte) {
	log.Println(fmt.Sprintf("client-%d: %s", ctx.GetConnection().GetConnID(), msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("%d-server1: %s", ctx.GetConnection().GetConnID(), msg))); err != nil {
		return
	}
	ctx.FireRead(msg)
}

func (d *DefaultServerChannelHandler) PostHandle(ctx iface.Context, msg []byte) {

}
