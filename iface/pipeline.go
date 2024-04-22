package iface

type Pipeline interface {
	AddLast(handler ConnectionHandler) // 在调用链末尾添加一个连接处理器
	Handle(msg []byte)                 // 根据调用链处理消息

	GetConnection() Connection // 获取连接对象
}
