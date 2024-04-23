package example

import (
	"log"
	"testing"

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
	clientBootstrap.GetConn().AddCronFunc("@every 1s", func() {
		// 每次发送一个数据包
		if err := clientBootstrap.GetConn().Write([]byte("hello world")); err != nil {
			log.Println(err)
			return
		}
	})

	select {}
}

func TestStartClient2(t *testing.T) {
	clientBootstrap := bootstrap.NewClient()
	clientBootstrap.SetEncoder(encoder.NewDelimiterBasedFrameDecoder('\n'))

	if err := clientBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

}
