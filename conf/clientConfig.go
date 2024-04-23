package conf

import (
	"runtime"
)

type ClientConfig struct {
	Ip   string
	Port int

	MaxPackageSize int // 最大包大小

	WorkerPoolSize int // 工作池大小
	TaskQueueSize  int // 最大工作队列长度
}

func DefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		Ip:             "127.0.0.1",
		Port:           8888,
		MaxPackageSize: 1024,

		WorkerPoolSize: runtime.NumCPU(),
		TaskQueueSize:  1024,
	}
}
