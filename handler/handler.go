package handler

import "context"

type ChannelHandler interface {
	ChannelRead(ctx context.Context, msg []byte)
}

type ConnectionHandler interface {
	ConnectionRead(ctx context.Context, msg []byte)
}
