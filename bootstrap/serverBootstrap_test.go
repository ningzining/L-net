package bootstrap

import (
	"github.com/ningzining/lazynet/iface"
	"log"
	"testing"
)

func TestStart(t *testing.T) {
	serverBootstrap := NewServerBootstrap(WithPort(8999))
	serverBootstrap.SetConnOnActiveFunc(func(conn iface.Connection) {
		log.Printf("remoteAddr: %s, connection on active", conn.RemoteAddr())
	})
	serverBootstrap.SetConnOnCloseFunc(func(conn iface.Connection) {
		log.Printf("remoteAddr: %s, connection on close", conn.RemoteAddr())
	})

	t.Log("tcp server start success")
	if err := serverBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

}
