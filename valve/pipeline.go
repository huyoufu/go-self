package valve

import (
	"container/list"
	"github.com/huyoufu/go-self/router"
)

type Pipeline struct {
	valves list.List
}

func New() *Pipeline {
	p := &Pipeline{}
	p.valves.Init()
	return p
}
func (p *Pipeline) Start(ctx router.Context) {
	p.start0(ctx)
}
func (p *Pipeline) start0(ctx router.Context) {
	next := p.next()
	next.process(ctx)

}
func (p *Pipeline) next() Valve {
	back := p.valves.Back()
	return back.Value.(Valve)
}
func (p *Pipeline) First(valve Valve) {
	p.valves.PushFront(valve)
}
func (p *Pipeline) FirstPF(pf func(ctx router.Context) bool) {
	p.valves.PushFront(ProcessFunc(pf))
}
func (p *Pipeline) Last(valve Valve) {
	p.valves.PushBack(valve)
}
func (p *Pipeline) LastPF(pf func(ctx router.Context) bool) {
	p.valves.PushBack(ProcessFunc(pf))
}
