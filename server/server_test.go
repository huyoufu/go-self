package server

import (
	"github.com/huyoufu/go-self/router"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer()
	server.EnableCors()
	server.EnableSession()
	router.Any("/x", func(ctx router.Context) {
		ctx.WriteString("hello world!!!")
	})
	server.Start()
}
