package self

import (
	"fmt"
	"strconv"
	"testing"
)

type XValve struct {
	name string
}

func newX(name string) *XValve {
	return &XValve{name}
}

func (x *XValve) process(ctx Context) {
	fmt.Println(1)
}

func TestPipeline_First(t *testing.T) {
	pl := NewPipeline()
	for i := 0; i < 100; i++ {

		pl.Last(newX("第" + strconv.Itoa(i) + "个"))
	}

	for _, v := range pl.valves {
		fmt.Println(v)
	}

}
func TestPipeline_Two(t *testing.T) {

}
