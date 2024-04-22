package iface

type Dispatcher interface {
	Dispatch(conn Connection, msg []byte)
	StartWorkerPool()
}
