package lazynet

import (
	"github.com/ningzining/lazynet/decoder"
	"testing"
)

func TestStart(t *testing.T) {
	serverBootstrap := NewServerBootstrap("0.0.0.0:9999").
		RegisterDecoder(decoder.NewDelimiterBasedFrameDecoder('\n'))

	t.Log("tcp server start success")
	if err := serverBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

}
