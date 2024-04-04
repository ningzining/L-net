package handler

import (
	"context"
	"log"
)

type DefaultConnectionHandler struct {
}

func NewDefaultConnectionHandler() *DefaultConnectionHandler {
	return &DefaultConnectionHandler{}
}

func (d *DefaultConnectionHandler) ConnectionRead(ctx context.Context, msg []byte) {
	log.Println("DefaultConnectionHandler.ConnectionRead start...")

	log.Println(string(msg))

	log.Println("DefaultConnectionHandler.ConnectionRead end...")
}
