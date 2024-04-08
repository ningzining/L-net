package iface

import "context"

type ChannelHandler interface {
	ChannelRead(ctx context.Context, msg []byte)
}

type ConnectionHandler interface {
	PreHandle(ctx Context, msg []byte)
	ConnectionRead(ctx Context, msg []byte)
	PostHandle(ctx Context, msg []byte)
}
