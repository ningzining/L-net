package bootstrap

import (
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"

	"net"
)

type ClientBootstrap struct {
	addr        string
	conn        net.Conn
	encoder     encoder.Encoder
	handlerList []iface.ChannelHandler2
}

func NewClientBootstrap(addr string) *ClientBootstrap {
	return &ClientBootstrap{
		addr: addr,
	}
}

// SetEncoder 设置编码器
func (c *ClientBootstrap) SetEncoder(encoder encoder.Encoder) {
	c.encoder = encoder
}

// GetEncoder 获取编码器
func (c *ClientBootstrap) GetEncoder() encoder.Encoder {
	return c.encoder
}

// Start 启动客户端
func (c *ClientBootstrap) Start() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *ClientBootstrap) Stop() {
	c.conn.Close()
}

func (c *ClientBootstrap) Read() ([]byte, error) {
	bytes := make([]byte, 1024)
	n, err := c.conn.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes[:n], nil
}

// 往连接中写入字节数组
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
