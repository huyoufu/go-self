package router

import (
	"container/list"
)

type Pipeline struct {
	valves list.List
}

func New() *Pipeline {
	p := &Pipeline{}
	p.valves.Init()
	return p
}
func (p *Pipeline) Invoke(ctx Context) {
	p.invoke0(ctx)
}
func (p *Pipeline) invoke0(ctx Context) {
	for item := p.valves.Front(); nil != item; item = item.Next() {
		v := item.Value.(Valve)
		if !v.process(ctx) {
			break
		}
	}
}

func (p *Pipeline) First(valve Valve) {
	p.valves.PushFront(valve)
}
func (p *Pipeline) FirstPF(vf ValveFunc) {
	p.valves.PushFront(vf)
}
func (p *Pipeline) Last(valve Valve) {
	p.valves.PushBack(valve)
}
func (p *Pipeline) LastPF(vf ValveFunc) {
	p.valves.PushBack(vf)
}
