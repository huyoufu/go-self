package self

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	app := NewServer()
	app.EnableCors()
	app.EnableSession()
	Any("/x", func(ctx Context) {
		ctx.WriteString("hello x!!!")
	})
	Any("/x/:id", func(ctx Context) {
		fmt.Println(ctx.PathParamValue("id"))

		ctx.WriteString("hello id!!!")
	})
	app.Start()
}
