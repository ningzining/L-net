package iface

import (
	"net"

	"github.com/robfig/cron/v3"
)

type Connection interface {
	GetConn() net.Conn // 获取连接对象
	GetConnID() uint32 // 获取连接id

	RemoteAddr() net.Addr // 获取远程地址

	Start() // 启动连接，开启读写操作
	Stop()  // 关闭连接，回收资源

	GetPipeline() Pipeline  // 获取当前连接的管道
	Write(msg []byte) error // 向当前连接写入数据

	AddCronFunc(spec string, cmd func()) error
	RemoveCronFunc(id cron.EntryID)
}
