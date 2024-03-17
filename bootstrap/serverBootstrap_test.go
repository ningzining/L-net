package bootstrap

import (
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/handler"
	"testing"
)

func TestStart(t *testing.T) {
	serverBootstrap := NewServerBootstrap("0.0.0.0:9999").
		RegisterDecoder(decoder.NewDelimiterBasedFrameDecoder('\n')).
		AddHandler(handler.NewChannelHandlerAdapter())

	t.Log("tcp server start success")
	if err := serverBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

}
