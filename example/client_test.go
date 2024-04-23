package example

import (
	"log"
	"testing"
	"time"

	"github.com/ningzining/lazynet/bootstrap"
	"github.com/ningzining/lazynet/conf"
	"github.com/ningzining/lazynet/decoder"
	"github.com/ningzining/lazynet/encoder"
)

func TestStartClient1(t *testing.T) {
	clientBootstrap := bootstrap.NewClient(conf.WithClientPort(8999))
	clientBootstrap.SetEncoder(encoder.NewLineBasedFrameDecoder())
	clientBootstrap.SetDecoder(decoder.NewLineBasedFrameDecoder())
	clientBootstrap.AddChannelHandler(NewDefaultClientChannelHandler())

	if err := clientBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}
	for true {
		// 每次发送一个数据包
		if err := clientBootstrap.GetConn().Write([]byte("hello world2")); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second * 1)
	}
}

func TestStartClient2(t *testing.T) {
	clientBootstrap := bootstrap.NewClient()
	clientBootstrap.SetEncoder(encoder.NewDelimiterBasedFrameDecoder('\n'))

	if err := clientBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

}
