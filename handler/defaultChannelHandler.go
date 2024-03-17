package handler

import (
	"context"
	log "github.com/ningzining/L-log"
)

type ChannelHandlerAdapter struct {
}

func NewChannelHandlerAdapter() ChannelHandler {
	return &ChannelHandlerAdapter{}
}

func (c ChannelHandlerAdapter) ChannelRead(ctx context.Context, msg []byte) {
	ctx = context.WithValue(ctx, "msg", string(msg))
	log.Info(string(msg))
}
