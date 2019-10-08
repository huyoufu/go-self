package valve

import (
	"container/list"
	"github.com/huyoufu/go-self/router"
)

type Pipeline struct {
	valves  list.List
	handler router.RouterHandler
}

func New() *Pipeline {
	p := &Pipeline{}
	p.valves.Init()
	return p
}
func (p *Pipeline) Invoke(ctx router.Context) {
	p.invoke0(ctx)
}
func (p *Pipeline) invoke0(ctx router.Context) {
	next := p.next()
	if next != nil {
		//如果还有下个valve 继续执行
		b := next.process(ctx)
		if b {
			p.invoke0(ctx)
		}
	} else {
		//否则结束 执行 handler
		p.handler.Service(ctx)

	}

}
func (p *Pipeline) next() Valve {
	back := p.valves.Back()
	return back.Value.(Valve)
}
func (p *Pipeline) First(valve Valve) {
	p.valves.PushFront(valve)
}
func (p *Pipeline) FirstPF(pf func(ctx router.Context) bool) {
	p.valves.PushFront(processFunc(pf))
}
func (p *Pipeline) Last(valve Valve) {
	p.valves.PushBack(valve)
}
func (p *Pipeline) LastPF(pf func(ctx router.Context) bool) {
	p.valves.PushBack(processFunc(pf))
}
