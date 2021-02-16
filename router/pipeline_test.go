package router

import (
	"fmt"
	"testing"
)

type X struct {
}

func (X) Foo(ctx *Context) bool {

	fmt.Println(1)
	return true
}

func TestPipeline_First(t *testing.T) {
	//pl:=NewPipeline()
	//pl.LastPF()
}
func TestPipeline_Two(t *testing.T) {

}
