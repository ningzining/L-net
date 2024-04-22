package handler

import (
	"fmt"
	"log"

	"github.com/ningzining/lazynet/iface"
)

type DefaultChannelHandler struct {
	BaseChannelHandler
}

func NewDefaultChannelHandler() iface.ChannelHandler {
	return &DefaultChannelHandler{}
}

func (d *DefaultChannelHandler) PreHandle(ctx iface.Context, msg []byte) {

}

func (d *DefaultChannelHandler) ChannelRead(ctx iface.Context, msg []byte) {
	log.Println(fmt.Sprintf("client-%d: %s", ctx.GetConnection().GetConnID(), msg))
	if err := ctx.GetConnection().Write([]byte(fmt.Sprintf("%d-server1: %s", ctx.GetConnection().GetConnID(), msg))); err != nil {
		return
	}
	ctx.FireRead(msg)
}

func (d *DefaultChannelHandler) PostHandle(ctx iface.Context, msg []byte) {

}
