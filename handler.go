package self

type Handler interface {
	Service(ctx Context)
}
type HandlerFunc func(ctx Context)

func (f HandlerFunc) Service(ctx Context) {
	f(ctx)
}
