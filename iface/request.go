package iface

type Request interface {
	GetConn() Connection // 获取本次请求的连接对象
	GetMsg() []byte      // 获取本次请求的消息
}
