package example

import (
	"testing"
	"time"

	"github.com/ningzining/lazynet/bootstrap"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
)

func TestStartClient1(t *testing.T) {
	clientBootstrap := bootstrap.NewClient("127.0.0.1:8999")
	clientBootstrap.SetEncoder(encoder.NewLineBasedFrameDecoder())
	clientBootstrap.SetDecoder(decoder.NewLineBasedFrameDecoder())
	clientBootstrap.AddChannelHandler(NewDefaultClientChannelHandler())

	if err := clientBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

	// 每次发送一个数据包
	for {
		if err := clientBootstrap.Write([]byte("hello world2")); err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Second * 2)
	}

}

func TestStartClient2(t *testing.T) {
	clientBootstrap := bootstrap.NewClient("127.0.0.1:8999")
	clientBootstrap.SetEncoder(encoder.NewDelimiterBasedFrameDecoder('\n'))

	if err := clientBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

	// 每次发送多个数据包
	for {
		if err := clientBootstrap.Write([]byte("hello world\nhello world\nhello world\nhello world\nhello world\nhello world")); err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Second)
	}

}
