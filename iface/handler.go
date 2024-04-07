package iface

import "context"

type ChannelHandler interface {
	ChannelRead(ctx context.Context, msg []byte)
}

type ConnectionHandler interface {
	ConnectionRead(ctx Context, msg []byte)
}
