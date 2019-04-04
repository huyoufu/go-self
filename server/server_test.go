package server

import (
	"github.com/huyoufu/go-self/router"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer()
	server.EnableCors()
	server.EnableSession()
	router.Any("/x/y/z/:id/:xx", func(ctx router.Context) {
		ctx.WriteString("你好")
	})
	server.Start()
}
