package bootstrap

import (
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/handler"
	"net"
)

type ClientBootstrap struct {
	addr        string
	conn        net.Conn
	encoder     encoder.Encoder
	handlerList []handler.ChannelHandler
}

func NewClientBootstrap(addr string) *ClientBootstrap {
	return &ClientBootstrap{
		addr: addr,
	}
}

func (c *ClientBootstrap) RegisterEncoder(encoder encoder.Encoder) *ClientBootstrap {
	c.encoder = encoder
	return c
}

func (c *ClientBootstrap) Start() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *ClientBootstrap) Write(source []byte) error {
	frame := source

	// 如果编码器不为nil，则对数据进行编码后写入
	var err error
	if c.encoder != nil {
		frame, err = c.encoder.Encode(frame)
		if err != nil {
			return err
		}
	}

	if _, err := c.conn.Write(frame); err != nil {
		return err
	}

	return nil
}
