package iface

type Pipeline interface {
	AddLast(handler ConnectionHandler)
	Handle(msg []byte)

	GetConnection() Connection
	SetConnection(conn Connection)
}
