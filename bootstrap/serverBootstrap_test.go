package bootstrap

import (
	"log"
	"testing"

	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/handler"
	"github.com/ningzining/lazynet/iface"
)

func TestStart1(t *testing.T) {
	serverBootstrap := NewServerBootstrap(WithPort(8999))
	serverBootstrap.SetConnOnActiveFunc(func(conn iface.Connection) {
		log.Printf("remoteAddr: %s, connection on active", conn.RemoteAddr())
	})
	serverBootstrap.SetConnOnCloseFunc(func(conn iface.Connection) {
		log.Printf("remoteAddr: %s, connection on close", conn.RemoteAddr())
	})
	serverBootstrap.SetDecoder(decoder.NewLineBasedFrameDecoder())
	serverBootstrap.SetEncoder(encoder.NewLineBasedFrameDecoder())
	serverBootstrap.AddConnectionHandler(handler.NewDefaultConnectionHandler())
	serverBootstrap.AddConnectionHandler(handler.NewDefaultConnectionHandler2())
	//serverBootstrap.AddConnectionHandler(handler.NewDefaultConnectionHandler3())

	if err := serverBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}
}

func TestStart2(t *testing.T) {
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
