package handler

import "context"

type ChannelHandler interface {
	ChannelRead(ctx context.Context, msg []byte)
}
