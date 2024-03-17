package bootstrap

import "testing"

func TestStartClientBootstrap(t *testing.T) {
	if err := NewClientBootstrap("127.0.0.1:9999").Start(); err != nil {
		t.Error(err)
		return
	}
}
