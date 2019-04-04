package router

import (
	"fmt"
)

type RouterHandler interface {
	Service(ctx Context)
}
type HandlerFunc func(ctx Context)

func (f HandlerFunc) Service(ctx Context) {
	f(ctx)
}

type Handler struct {
	Status int
	Name   string
}

func (h *Handler) Service(ctx Context) {
	fmt.Println(h.Status, h.Name)
	fmt.Println("handler 正在处理请求")
}

var NotFoundHandler RouterHandler = &Handler{
	404,
	"not found",
}
