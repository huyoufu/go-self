package valve

import "github.com/huyoufu/go-self/router"

type Valve interface {
	process(ctx router.Context) bool
}
type processFunc func(ctx router.Context) bool

func (f processFunc) process(ctx router.Context) bool {
	return f(ctx)
}
