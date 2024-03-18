package bootstrap

import (
	"bytes"
	"context"
	"errors"
	log "github.com/ningzining/L-log"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/handler"
	"net"
)

type ServerBootstrap struct {
	addr        string
	decoder     decoder.Decoder
	handlerList []handler.ChannelHandler
}

func NewServerBootstrap(addr string) *ServerBootstrap {
	return &ServerBootstrap{
		addr:        addr,
		handlerList: make([]handler.ChannelHandler, 0),
	}
}

// RegisterDecoder 注册解码器
func (s *ServerBootstrap) RegisterDecoder(d decoder.Decoder) *ServerBootstrap {
	s.decoder = d
	return s
}

// AddHandler 添加处理器
func (s *ServerBootstrap) AddHandler(handler handler.ChannelHandler) *ServerBootstrap {
	s.handlerList = append(s.handlerList, handler)
	return s
}

// Start 启动服务端
func (s *ServerBootstrap) Start() error {
	// 校验相关参数是否正常
	if err := s.verifyParam(); err != nil {
		return err
	}

	// 监听端口
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err.Error())
			continue
		}
		go s.handleConn(conn)
	}
}

// 验证参数是否完整
func (s *ServerBootstrap) verifyParam() error {
	if s.addr == "" {
		return errors.New("addr must be required")
	}
	if s.decoder == nil {
		return errors.New("decoder must be required")
	}
	if len(s.handlerList) == 0 {
		return errors.New("handler must be required")
	}

	return nil
}

// 处理连接
func (s *ServerBootstrap) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()

		if err := recover(); err != nil {
			log.Errorf("%v", err)
		}
	}()

	// 新建缓冲区
	var buffer = bytes.NewBuffer(make([]byte, 0, 4096))
	// 初始化上下文
	ctx := handler.NewChannelHandlerContext(context.Background())

	for {
		// 读取数据
		readBytes := make([]byte, 1024)
		n, err := conn.Read(readBytes)
		if err != nil {
			break
		}

		// 写入buffer缓冲区
		_, err = buffer.Write(readBytes[:n])
		if err != nil {
			log.Errorf("%v", err)
			continue
		}

		// 处理缓冲区数据
		for {
			// 使用预加载的解码器解析数据包供下方处理器使用
			msg, err := s.decoder.Decode(buffer)
			if err != nil {
				break
			}
			// 对每一个数据包做处理
			for _, channelHandler := range s.handlerList {
				channelHandler.ChannelRead(ctx, msg)
			}
		}
	}
}
