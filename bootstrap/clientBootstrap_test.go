package bootstrap

import (
	"github.com/ningzining/lazynet/encoder"
	"testing"
	"time"
)

func TestStartClientBootstrap(t *testing.T) {
	clientBootstrap := NewClientBootstrap("127.0.0.1:8999").
		RegisterEncoder(encoder.NewDelimiterBasedFrameDecoder('\n'))

	if err := clientBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

	for {
		if err := clientBootstrap.Write([]byte("hello world")); err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Second)
	}

}
