package pipeline

import (
	"context"

	"github.com/ningzining/lazynet/iface"
)

type Pipeline struct {
	head *ConnectionContext
	tail *ConnectionContext

	connection iface.Connection
}

func NewPipeline(conn iface.Connection) iface.Pipeline {
	p := &Pipeline{
		connection: conn,
	}
	head := NewContext(context.Background(), p, nil)
	tail := NewContext(context.Background(), p, nil)

	head.next = tail
	tail.prev = head

	p.head = head
	p.tail = tail

	return p
}

func (p *Pipeline) AddLast(handler iface.ConnectionHandler) {
	ctx := NewContext(context.Background(), p, handler)
	prev := p.tail.prev
	ctx.prev = prev
	ctx.next = p.tail
	prev.next = ctx
	p.tail.prev = ctx
}

func (p *Pipeline) Handle(msg []byte) {
	ctx := p.firstContext()

	ctx.DoHandle(msg)
}

func (p *Pipeline) firstContext() iface.Context {
	node := p.head
	for {
		if node != nil && node.handler != nil {
			return node
		}

		node = node.next
	}
}

func (p *Pipeline) GetConnection() iface.Connection {
	return p.connection
}

func (p *Pipeline) SetConnection(conn iface.Connection) {
	p.connection = conn
}
