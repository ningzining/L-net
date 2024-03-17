package bootstrap

import (
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/handler"
	"net"
	"time"
)

type ClientBootstrap struct {
	addr        string
	decoder     decoder.Decoder
	handlerList []handler.ChannelHandler
}

func NewClientBootstrap(addr string) *ClientBootstrap {
	return &ClientBootstrap{
		addr: addr,
	}
}

func (c *ClientBootstrap) Start() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}

	for {
		_, err = conn.Write([]byte("hello\n"))
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}
