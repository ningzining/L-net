package conf

import (
	"runtime"
)

type Config struct {
	Host string // IP地址
	Port int    // 端口

	MaxPackageSize int // 最大包大小
	MaxConnSize    int // 最大连接数

	WorkerPoolSize int // 工作池大小
	TaskQueueSize  int // 最大工作队列长度
}

func DefaultConfig() *Config {
	return &Config{
		Host:           "0.0.0.0",
		Port:           8888,
		MaxPackageSize: 1024,
		MaxConnSize:    1024 * 10,
		WorkerPoolSize: runtime.NumCPU(),
		TaskQueueSize:  1024,
	}
}
