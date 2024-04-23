package conf

type ClientOption func(c *ClientConfig)

func WithClientPort(port int) ClientOption {
	return func(c *ClientConfig) {
		c.Port = port
	}
}
