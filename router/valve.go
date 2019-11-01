package router

type Valve interface {
	process(ctx Context)
}
type ValveFunc func(ctx Context)

func (f ValveFunc) process(ctx Context) {
	f(ctx)
}
