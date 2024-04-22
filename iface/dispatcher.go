package iface

type Dispatcher interface {
	StartWorkerPool()     // 开启工作协程池
	Dispatch(req Request) // 分发请求到协程池中
}
