package connection

import (
	"bytes"
	"net"

	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
	"github.com/ningzining/lazynet/iface"
	"github.com/ningzining/lazynet/server"
)

type Connection struct {
	bootstrap iface.Bootstrap

	conn       net.Conn
	connId     uint32
	remoteAddr net.Addr
	localAddr  net.Addr

	decoder decoder.Decoder // 解码器
	encoder encoder.Encoder // 编码器

	pipeline iface.Pipeline // 处理器管道

	readBuffer *bytes.Buffer // 读取缓冲区

	exitChan chan struct{} // 退出管道
	msgChan  chan []byte   // 读写goroutine管道

	onActive func(conn iface.Connection) // 钩子函数，当连接建立的时候调用
	onClose  func(conn iface.Connection) // 钩子函数，当连接断开的时候调用
}

func New(bootstrap iface.Bootstrap, conn net.Conn, connId uint32, maxPackageSize int) iface.Connection {
	c := &Connection{
		bootstrap:  bootstrap,
		conn:       conn,
		connId:     connId,
		remoteAddr: conn.RemoteAddr(),
		localAddr:  conn.LocalAddr(),
		decoder:    bootstrap.GetDecoder(),
		encoder:    bootstrap.GetEncoder(),
		pipeline:   nil,
		readBuffer: bytes.NewBuffer(make([]byte, 0, maxPackageSize*2)), // 缓冲区大小为最大包大小的两倍，确保能解析出来一个包
		msgChan:    make(chan []byte),
		exitChan:   make(chan struct{}),
		onActive:   nil,
		onClose:    nil,
	}
	pipeline := server.NewPipeline(c)
	c.pipeline = pipeline

	for _, handler := range bootstrap.GetChannelHandlers() {
		c.pipeline.AddLast(handler)
	}

	return c
}

func (c *Connection) GetConn() net.Conn {
	return c.conn
}

func (c *Connection) GetConnID() uint32 {
	return c.connId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.remoteAddr
}

func (c *Connection) Start() {
	// 执行连接建立的钩子函数
	c.callOnActive()

	go c.StartReader()
	// 启动写入器
	go c.StartWriter()
}

func (c *Connection) callOnActive() {
	if c.onActive != nil {
		c.onActive(c)
	}
}

func (c *Connection) StartReader() {
	defer func() {
		c.Stop()
		if err := recover(); err != nil {
			log.Errorf("%v", err)
		}
	}()

	defer log.Infof("%s: [reader] end", c.RemoteAddr().String())
	log.Infof("%s: [reader] start", c.RemoteAddr().String())

	for {
		// 读取数据
		readBytes := make([]byte, 1024)
		n, err := c.conn.Read(readBytes)
		if err != nil {
			break
		}
		// 写入连接的缓冲区
		c.readBuffer.Write(readBytes[:n])
		// 使用注册的解码器进行解码
		if c.decoder != nil {
			// 一个数据包可能包含多个数据帧的情况，所以需要循环处理
			frames := c.decoder.Decode(c.readBuffer)
			// 读取每一帧的数据并进行处理
			for _, frame := range frames {
				// 创建一个请求体
				req := server.NewRequest(c, frame)
				// 使用消息分发器分发消息，异步处理请求
				go c.bootstrap.GetDispatcher().Dispatch(req)
			}
		} else {
			// 创建一个请求体
			req := server.NewRequest(c, c.readBuffer.Bytes())
			// 使用消息分发器分发消息，异步处理请求
			go c.bootstrap.GetDispatcher().Dispatch(req)
			// 处理完消息，重置缓冲区
			c.readBuffer.Reset()
		}
	}
}

// StartWriter
// 实现写入器
func (c *Connection) StartWriter() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("%v", err)
		}
	}()
	defer log.Infof("%s: [writer] end", c.RemoteAddr().String())
	log.Infof("%s: [writer] start", c.RemoteAddr().String())

	// 等待监听写入的消息
	for {
		select {
		case data := <-c.msgChan:
			// 使用编码器对返回的数据进行编码
			if c.encoder != nil {
				var err error
				if data, err = c.encoder.Encode(data); err != nil {
					return
				}
			}

			// 向客户端写入数据
			if _, err := c.conn.Write(data); err != nil {
				return
			}
		case <-c.exitChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	// 执行连接关闭的钩子函数
	c.callOnClose()

	// 关闭连接
	c.conn.Close()

	// 告知退出
	c.exitChan <- struct{}{}

	// 回收资源
	close(c.exitChan)
	close(c.msgChan)
}

func (c *Connection) callOnClose() {
	if c.onClose != nil {
		c.onClose(c)
	}
}

func (c *Connection) GetPipeline() iface.Pipeline {
	return c.pipeline
}

func (c *Connection) Write(msg []byte) error {
	frame := msg

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
