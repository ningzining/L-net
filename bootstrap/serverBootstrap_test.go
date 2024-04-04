package bootstrap

import (
	"testing"
)

func TestStart(t *testing.T) {
	serverBootstrap := NewServerBootstrap(WithPort(8999))

	t.Log("tcp server start success")
	if err := serverBootstrap.Start(); err != nil {
		t.Error(err)
		return
	}

}
