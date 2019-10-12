package router

type Valve interface {
	process(ctx Context) bool
}
type ValveFunc func(ctx Context) bool

func (f ValveFunc) process(ctx Context) bool {
	return f(ctx)
}
