package iface

import (
	"context"
)

type Context interface {
	context.Context
	GetConnection() Connection // 根据上下文获取当前的连接
}
