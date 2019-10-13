package router

import (
	"fmt"
	"testing"
)

type X struct {
}

func (X) Haha(ctx *Context) bool {

	fmt.Println(1)
	return true
}

func TestPipeline_First(t *testing.T) {

}
func TestPipeline_Two(t *testing.T) {

	fmt.Println(1)
}
