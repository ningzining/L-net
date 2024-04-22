package iface

type ConnManager interface {
	Add(conn Connection)  // 添加连接
	Remove(connId uint32) // 移除连接

	Get(connId uint32) (Connection, bool) // 根据连接id查询连接
	Size() int                            // 获取连接数量

	Clear() // 清理所有连接
}
