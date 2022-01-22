package router

type HandlerPipeline struct {
	Handler
	pl *Pipeline
}
type Pipeline struct {
	pos, n int
	valves []Valve
}

func NewPipeline() *Pipeline {
	p := &Pipeline{}
	p.valves = nil
	return p
}

func NewRouterHandlerPipeline(pl *Pipeline) *HandlerPipeline {
	hp := &HandlerPipeline{}
	hp.pl = pl
	return hp
}
func Compose(spl, hpl *Pipeline) *Pipeline {
	result := NewPipeline()
	for _, v := range spl.valves {
		if v != nil {
			result.valves = append(result.valves, v)
		}
	}
	for _, v := range hpl.valves {
		if v != nil {
			result.valves = append(result.valves, v)
		}
	}
	return result

}
func (hp *HandlerPipeline) GetPipeLine() *Pipeline {
	return hp.pl
}
func (hp *HandlerPipeline) Invoke(ctx Context) {
	hp.invoke0(ctx)
}
func (hp *HandlerPipeline) invoke0(ctx Context) {
	/*for item := hp.pl.valves.Front(); nil != item; item = item.Next() {
		v := item.Value.(Valve)
		v.process(ctx)
	}*/
	pos := hp.pl.pos
	n := len(hp.pl.valves)
	if pos < n {
		ok, pos, v := mustValve(pos, hp.pl.valves)
		if ok {
			hp.pl.pos = pos
			v.process(ctx)
		} else {
			hp.Handler.Service(ctx)
		}

	} else {
		hp.Handler.Service(ctx)
	}

}

func (p *Pipeline) Last(valve Valve) *Pipeline {
	p.valves = append(p.valves, valve)

	return p
}
func (p *Pipeline) LastPF(vf ValveFunc) *Pipeline {
	p.valves = append(p.valves, vf)

	return p

}
func mustValve(start int, valves []Valve) (bool, int, Valve) {
	for ; start < len(valves); start++ {
		valve := valves[start]
		if valve != nil {
			return true, start + 1, valve
		}
	}
	return false, 0, nil
}
