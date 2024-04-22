package example

import (
	"fmt"
	"log"

	"github.com/ningzining/lazynet/handler"
	"github.com/ningzining/lazynet/iface"
)

type DefaultClientChannelHandler struct {
	handler.BaseChannelHandler
}

func NewDefaultClientChannelHandler() iface.ChannelHandler {
	return &DefaultClientChannelHandler{}
}

func (d *DefaultClientChannelHandler) PreHandle(ctx iface.Context, msg []byte) {

}

func (d *DefaultClientChannelHandler) ChannelRead(ctx iface.Context, msg []byte) {
	log.Println(fmt.Sprintf("client-%d: %s", ctx.GetConnection().GetConnID(), msg))
	ctx.FireRead(msg)
}

func (d *DefaultClientChannelHandler) PostHandle(ctx iface.Context, msg []byte) {

}
