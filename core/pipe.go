package core

import "github.com/tk103331/logpipe/config"

type Pipe struct {
	conf    PipeConf
	Ctx     Context
	Inputs  []Input
	Filters []Filter
	Outputs []Output
}

func NewPipe(name string, value *config.Value) *Pipe {
	conf := PipeConf{}
	conf.Load(value)
	return &Pipe{conf: conf}
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
	Inputs  []*config.BaseConf
	Filters []*config.BaseConf
	Outputs []*config.BaseConf
}

func (c *PipeConf) Load(value *config.Value) {
	_ = c.BaseConf.Load(value)
	c.Inputs = c.GetArray("input")
	c.Filters = c.GetArray("filter")
	c.Outputs = c.GetArray("output")
}
