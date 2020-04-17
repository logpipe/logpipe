package core

import "github.com/tk103331/logpipe/config"

type Pipe struct {
	Ctx     Context
	Inputs  []Input
	Filters []Filter
	Outputs []Output
}

func NewPipe(name string, value config.Value) *Pipe {

}

func (p *Pipe) Start() {
	for _, output := range p.Outputs {
		output.Start()
	}
	for _, input := range p.Inputs {
		input.Start(p.Ctx)
	}
}

func (p *Pipe) Stop() {

	for _, input := range p.Inputs {
		input.Stop(p.Ctx)
	}

	for _, output := range p.Outputs {
		output.Stop()
	}
}

type PipeConf struct {
	config.BaseConf
	Inputs  []*config.Value
	Filters []*config.Value
	Outputs []*config.Value
}

func (c *PipeConf) Load(value config.Value) {
	c.BaseConf.Load(value)
	c.Inputs = value.GetArray("input")
	c.Inputs = value.GetArray("filter")
	c.Outputs = value.GetArray("output")
}
