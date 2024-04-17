package iface

import (
	"context"
)

type Context interface {
	context.Context

	GetHandler() ConnectionHandler
	GetConnection() Connection

	DoHandle(msg []byte)
	FireConnectionRead(msg []byte)
}
