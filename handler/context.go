package handler

import "context"

type ChannelHandlerContext struct {
	context.Context
}

func NewChannelHandlerContext(ctx context.Context) context.Context {
	return &ChannelHandlerContext{Context: ctx}
}
