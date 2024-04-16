package pipeline

import (
	"github.com/ningzining/lazynet/iface"
)

type Pipeline struct {
	head *Node
	tail *Node
}

func NewPipeline() iface.Pipeline {
	head := NewNode(nil)
	tail := NewNode(nil)
	head.next = tail
	tail.prev = head

	return &Pipeline{
		head: head,
		tail: tail,
	}
}

func (p *Pipeline) AddLast(handler iface.ConnectionHandler) {
	node := NewNode(handler)
	p.tail.prev.next = node
	node.prev = p.tail.prev

	p.tail.prev = node
	node.next = p.tail
}

func (p *Pipeline) Handle(ctx iface.Context, msg []byte) {
	node := p.head
	for {
		if node.next == nil || node.next.handler == nil {
			return
		}

		node.next.handler.PreHandle(ctx, msg)
		node.next.handler.ConnectionRead(ctx, msg)
		node.next.handler.PostHandle(ctx, msg)

		node = node.next
	}
}
