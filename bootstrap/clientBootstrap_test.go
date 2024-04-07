package bootstrap

import (
	"testing"
	"time"

	"github.com/ningzining/lazynet/encoder"
)

func TestStartClientBootstrap1(t *testing.T) {
	clientBootstrap := NewClientBootstrap("127.0.0.1:8999").
		RegisterEncoder(encoder.NewDelimiterBasedFrameDecoder('\n'))

	if err := clientBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

	// 每次发送一个数据包
	for {
		if err := clientBootstrap.Write([]byte("hello world")); err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Second)
	}

}

func TestStartClientBootstrap2(t *testing.T) {
	clientBootstrap := NewClientBootstrap("127.0.0.1:8999").
		RegisterEncoder(encoder.NewDelimiterBasedFrameDecoder('\n'))

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
