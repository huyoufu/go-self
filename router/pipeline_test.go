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

	valves := make([]Valve, 8)
	for v := range valves {

		fmt.Println(v)
	}

}
