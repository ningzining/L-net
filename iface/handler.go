package iface

type ChannelHandler interface {
	PreHandle(ctx Context, msg []byte)   // 预处理
	ChannelRead(ctx Context, msg []byte) // 处理每一帧的数据
	PostHandle(ctx Context, msg []byte)  // 后处理
}
