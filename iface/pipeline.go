package iface

type Pipeline interface {
	AddLast(handler ConnectionHandler)
	Handle(ctx Context, msg []byte)
}
