package handler

type ChannelHandler interface {
	ChannelRead(msg []byte)
}
