package router

type RouterHandler interface {
	Service(ctx Context)
}
type HandlerFunc func(ctx Context)

func (f HandlerFunc) Service(ctx Context) {
	f(ctx)
}

type RouterHandlerPipeline struct {
	RouterHandler
	pl *Pipeline
}
