package pipeline

import (
	"github.com/ningzining/lazynet/iface"
)

type Node struct {
	handler iface.ConnectionHandler

	prev *Node
	next *Node
}

func NewNode(handler iface.ConnectionHandler) *Node {
	return &Node{
		handler: handler,
		prev:    nil,
		next:    nil,
	}
}

func (n *Node) AddNext(node *Node) {
	n.next = node
	node.prev = n
}
