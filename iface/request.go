package iface

type Request interface {
	GetConn() Connection
	GetMsg() []byte
}
