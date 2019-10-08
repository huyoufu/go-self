package valve

import (
	"fmt"
	"github.com/huyoufu/go-self/router"
	"testing"
)

type X struct {
}

func (X) Haha(ctx router.Context) bool {

	fmt.Println(1)
	return true
}

func TestPipeline_First(t *testing.T) {

	pipeline := Pipeline{}
	pipeline.FirstPF(X{}.Haha)
	pipeline.Start(nil)
}
func TestPipeline_Two(t *testing.T) {

	fmt.Println(1)
}
