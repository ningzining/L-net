package iface

type Server interface {
	Start() error
	Stop()
}
