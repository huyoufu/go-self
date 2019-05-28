package valve

import "github.com/huyoufu/go-self/router"

type Valve interface {
	process(ctx router.Context) bool
}
type ProcessFunc func(ctx router.Context) bool

func (f ProcessFunc) process(ctx router.Context) {
	f(ctx)
}
