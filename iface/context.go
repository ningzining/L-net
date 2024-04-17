package iface

import (
	"context"
)

type Context interface {
	context.Context

	GetConnection() Connection

	GetHandler() ConnectionHandler
	FireConnectionRead(msg []byte)
}
